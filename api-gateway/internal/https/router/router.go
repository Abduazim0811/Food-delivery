package router

import (
	deliveryclient "api-gateway/internal/clients/delivery-client"
	orderclient "api-gateway/internal/clients/order-client"
	paymentclients "api-gateway/internal/clients/payment-clients"
	productclients "api-gateway/internal/clients/product-clients"
	userclient "api-gateway/internal/clients/user-client"
	deliveryhandler "api-gateway/internal/https/handlers/delivery-service"
	orderhandler "api-gateway/internal/https/handlers/order-service"
	paymenthandler "api-gateway/internal/https/handlers/payment-service"
	producthandler "api-gateway/internal/https/handlers/product-handler"
	"api-gateway/internal/https/handlers/userhandler"
	"api-gateway/internal/pkg/jwt"
	middleware "api-gateway/internal/rate-limiting"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func Router() *http.Server {
	userClient := userclient.DialGrpcUser()
	userHandler := &userhandler.UserHandler{ClientUser: userClient}

	productClient := productclients.DialGrpcProduct()
	productHandler := &producthandler.ProductHandler{ClientProduct: productClient}

	paymentClient := paymentclients.DialGrpcPayment()
	paymentHandler := &paymenthandler.PaymentHandler{ClientPayment: paymentClient}

	orderClient := orderclient.DialGrpcOrder()
	orderHandler := &orderhandler.OrderHandler{ClientOrder: orderClient}

	deliveryClient := deliveryclient.DialGrpcDelivery()
	deliveryHandler := &deliveryhandler.DeliveryHandler{ClientDelivery: deliveryClient}

	userRateLimiter := middleware.NewRateLimiter(2, time.Minute)
	productRateLimiter := middleware.NewRateLimiter(3, time.Minute)
	paymentRateLimiter := middleware.NewRateLimiter(4, time.Minute)
	orderRateLimiter := middleware.NewRateLimiter(4, time.Minute)
	deliveryRateLimiter := middleware.NewRateLimiter(4, time.Minute)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRoutes := r.Group("/users")
	userRoutes.Use(userRateLimiter.Limit())
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/verify-code", userHandler.VerifyCode)
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.GET("/:id", jwt.TokenMiddleware("user"),userHandler.GetbyIdUser)
		userRoutes.PUT("/", jwt.TokenMiddleware("user"),userHandler.UpdateUser)
		userRoutes.DELETE("/:id", jwt.TokenMiddleware("user"), userHandler.DeleteUser)
	}

	courierRoutes := r.Group("/couriers")
	courierRoutes.Use(userRateLimiter.Limit())
	{
		courierRoutes.POST("/register", userHandler.RegisterCourier)
		courierRoutes.POST("/verify-code", userHandler.VerifyCodeCourier)
		courierRoutes.POST("/login", userHandler.LoginCourier)
		courierRoutes.PUT("/", jwt.TokenMiddleware("courier"), userHandler.UpdateCourier)
		courierRoutes.DELETE("/:id",jwt.TokenMiddleware("courier"), userHandler.DeleteCourier)
	}

	productRoutes := r.Group("/")
	productRoutes.Use(jwt.TokenMiddleware("user"))
	productRoutes.Use(productRateLimiter.Limit())
	{
		productRoutes.POST("/products", productHandler.CreateProduct)
		productRoutes.GET("/products/:id", productHandler.GetByIdProduct)
		productRoutes.GET("/products", productHandler.GetAllProducts)
		productRoutes.PUT("/products", productHandler.UpdateProduct)
		productRoutes.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	paymentRoutes := r.Group("/")
	paymentRoutes.Use(paymentRateLimiter.Limit())
	{
		paymentRoutes.POST("/payment/process", paymentHandler.ProcessPayment)
		paymentRoutes.POST("/payment/refund", paymentHandler.RefundPayment)
	}

	orderRoutes := r.Group("/")
	orderRoutes.Use(jwt.TokenMiddleware("courier"))
	orderRoutes.Use(orderRateLimiter.Limit())
	{
		orderRoutes.POST("/orders", orderHandler.CreateOrder)
		orderRoutes.GET("/orders/:id", orderHandler.GetbyIdOrder)
		orderRoutes.PUT("/orders", orderHandler.UpdateOrder)
		orderRoutes.DELETE("/orders/:id", orderHandler.DeleteOrder)
	}

	deliveryRoutes := r.Group("/")
	deliveryRoutes.Use(jwt.TokenMiddleware("courier"))
	deliveryRoutes.Use(deliveryRateLimiter.Limit())
	{
		deliveryRoutes.POST("/delivery", deliveryHandler.CreateDelivery)
		deliveryRoutes.GET("/delivery/:id", deliveryHandler.GetDeliveryStatus)
		deliveryRoutes.PUT("/delivery/update-status", deliveryHandler.UpdateDeliveryStatus)
	}
	server := &http.Server{
		Addr:    os.Getenv("server_url"),
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServeTLS("./internal/tls/items.pem", "./internal/tls/items-key.pem"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run HTTPS server: %v", err)
		}
	}()

	GracefulShutdown(server, log.Default())

	return server
}

func GracefulShutdown(srv *http.Server, logger *log.Logger) {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	<-shutdownCh
	logger.Println("Shutdown signal received, initiating graceful shutdown...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Printf("Server shutdown encountered an error: %v", err)
	} else {
		logger.Println("Server gracefully stopped")
	}

	select {
	case <-shutdownCtx.Done():
		if shutdownCtx.Err() == context.DeadlineExceeded {
			logger.Println("Shutdown deadline exceeded, forcing server to stop")
		}
	default:
		logger.Println("Shutdown completed within the timeout period")
	}
}
