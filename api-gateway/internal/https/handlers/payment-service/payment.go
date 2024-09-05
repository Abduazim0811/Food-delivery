package paymenthandler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"api-gateway/internal/protos/paymentproto"
)

type PaymentHandler struct {
	ClientPayment paymentproto.PaymentServiceClient
}

// @Summary Process a payment
// @Description Processes a payment based on the provided payment details
// @Tags payment
// @Accept json
// @Produce json
// @Param payment body paymentproto.ProcessPaymentRequest true "Payment processing request"
// @Success 200 {object} paymentproto.ProcessPaymentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /payment/process [post]
func (p *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req paymentproto.ProcessPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientPayment.ProcessPayment(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Refund a payment
// @Description Refunds a payment based on the provided refund details
// @Tags payment
// @Accept json
// @Produce json
// @Param refund body paymentproto.RefundPaymentRequest true "Payment refund request"
// @Success 200 {object} paymentproto.RefundPaymentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /payment/refund [post]
func (p *PaymentHandler) RefundPayment(c *gin.Context) {
	var req paymentproto.RefundPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientPayment.RefundPayment(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
