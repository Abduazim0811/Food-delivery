package deliveryhandler

import (
	"context"
	"net/http"
	"time"

	"api-gateway/internal/protos/deliveryproto"

	"github.com/gin-gonic/gin"
)

type DeliveryHandler struct {
	ClientDelivery deliveryproto.DeliveryServiceClient
}

// @Summary Create a new delivery
// @Description Creates a new delivery with the provided details
// @Tags delivery
// @Accept json
// @Produce json
// @Param delivery body deliveryproto.CreateDeliveryReq true "Delivery request body"
// @Success 200 {object} deliveryproto.CreateDeliveryRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /delivery [post]
func (d *DeliveryHandler) CreateDelivery(c *gin.Context) {
	var req deliveryproto.CreateDeliveryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := d.ClientDelivery.CreateDelivery(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get delivery status by ID
// @Description Retrieves the status of a delivery by its ID
// @Tags delivery
// @Accept json
// @Produce json
// @Param id path int true "Delivery ID"
// @Success 200 {object} deliveryproto.Delivery
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /delivery/{id} [get]
func (d *DeliveryHandler) GetDeliveryStatus(c *gin.Context) {
	id := c.Param("id")
	req := &deliveryproto.GetDeliveryStatusReq{DeliveryId: id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := d.ClientDelivery.GetDeliveryStatus(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update delivery status
// @Description Updates the status of an existing delivery
// @Tags delivery
// @Accept json
// @Produce json
// @Param delivery body deliveryproto.UpdateDeliveryStatusReq true "Delivery update status request"
// @Success 200 {object} deliveryproto.UpdateDeliveryStatusRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /delivery/update-status [put]
func (d *DeliveryHandler) UpdateDeliveryStatus(c *gin.Context) {
	var req deliveryproto.UpdateDeliveryStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := d.ClientDelivery.UpdateDeliveryStatus(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
