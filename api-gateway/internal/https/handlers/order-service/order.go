package orderhandler

import (
	"context"
	"net/http"
	"time"

	"api-gateway/internal/protos/orderproto"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	ClientOrder orderproto.OrderServiceClient
}

// @Summary Create a new order
// @Description Creates a new order with the provided details
// @Tags order
// @Accept json
// @Produce json
// @Param order body orderproto.CreateOrderReq true "Order request body"
// @Success 200 {object} orderproto.CreateOrderRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /orders [post]
func (o *OrderHandler) CreateOrder(c *gin.Context) {
	var req orderproto.CreateOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.CreateOrder(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get order by ID
// @Description Retrieves an order by its ID
// @Tags order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} orderproto.GetOrderRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /orders/{id} [get]
func (o *OrderHandler) GetbyIdOrder(c *gin.Context) {
	id := c.Param("id")
	req := &orderproto.GetOrderReq{OrderId: id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.GetbyIdOrder(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update an order
// @Description Updates an existing order with the provided details
// @Tags order
// @Accept json
// @Produce json
// @Param order body orderproto.UpdateReq true "Order update request"
// @Success 200 {object} orderproto.UpdateOrderRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /orders [put]
func (o *OrderHandler) UpdateOrder(c *gin.Context) {
	var req orderproto.UpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.UpdateOrder(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete an order
// @Description Deletes an order by its ID
// @Tags order
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} orderproto.UpdateOrderRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /orders/{id} [delete]
func (o *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	req := &orderproto.GetOrderReq{OrderId: id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.DeleteOrder(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
