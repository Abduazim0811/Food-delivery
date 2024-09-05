package producthandler

import (
	"api-gateway/internal/protos/productproto"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ClientProduct productproto.ProductServiceClient
}

// @Summary Create a new product
// @Description Create a new product with the given details
// @Tags product
// @Accept json
// @Produce json
// @Param product body productproto.CreateReq true "Product creation request"
// @Success 200 {object} productproto.CreateRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /products [post]
func (p *ProductHandler) CreateProduct(c *gin.Context) {
	var req productproto.CreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.CreateProduct(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get product by ID
// @Description Retrieve product details by product ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} productproto.Product
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /products/{id} [get]
func (p *ProductHandler) GetByIdProduct(c *gin.Context) {
	id := c.Param("id")
	req := &productproto.ProductResponse{Id: id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.GetByIdProduct(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} productproto.ListProduct
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /products [get]
func (p *ProductHandler) GetAllProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.GetAllProducts(ctx, &productproto.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update product details
// @Description Update the details of an existing product
// @Tags product
// @Accept json
// @Produce json
// @Param product body productproto.Product true "Product update request"
// @Success 200 {object} productproto.CreateRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /products [put]
func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	var req productproto.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.UpdateProduct(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete product by ID
// @Description Delete an existing product by product ID
// @Tags product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} productproto.CreateRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /products/{id} [delete]
func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	req := &productproto.ProductResponse{Id: id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.DeleteProduct(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
