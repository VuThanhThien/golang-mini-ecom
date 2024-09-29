package controllers

import (
	"net/http"
	"strconv"

	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/payment/internal/models/dto"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	service services.PaymentServiceInterface
}

func NewPaymentController(service services.PaymentServiceInterface) *PaymentController {
	return &PaymentController{service: service}
}

// CreatePayment godoc
//
// @Summary		CreatePayment
// @Description	CreatePayment
// @Tags			Payments
// @Accept			json
// @Produce		json
// @Param			payment	body		dto.CreatePaymentDto	true	"payment"
// @Security		Bearer
// @Router			/payments [post]
// @Success		201	{object}		object
func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	var input dto.CreatePaymentDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment, err := c.service.CreatePayment(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, payment)
}

// ListPayments godoc
//
// @Summary		ListPayments
// @Description	List payments with optional filtering and pagination
// @Tags			Payments
// @Accept			json
// @Produce		json
// @Param filter query dto.FilterPaymentDto false "Filter parameters"
// @Param _ query dto.PaginationDto false "PaginationDto"
// @Security		Bearer
// @Router			/payments [get]
// @Success		200	{object}	object
// @Failure		400	{object}	object
// @Failure		500	{object}	object
func (c *PaymentController) ListPayments(ctx *gin.Context) {
	var filter dto.FilterPaymentDto

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	payments, err := c.service.ListPayments(filter, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)
}

// ReadPayment godoc
//
// @Summary		ReadPayment
// @Description	Read a payment by ID
// @Tags			Payments
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Payment ID"
// @Security		Bearer
// @Router			/payments/{id} [get]
// @Success		200	{object}	object
// @Failure		400	{object}	object
// @Failure		500	{object}	object
func (c *PaymentController) ReadPayment(ctx *gin.Context) {
	var readIdRequest dto.ReadIdRequest
	if err := ctx.ShouldBindUri(&readIdRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payment, err := c.service.ReadPayment(uint(readIdRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payment)
}

// ReadByOrderId godoc
//
// @Summary		ReadByOrderId
// @Description	Read payments by order ID
// @Tags			Payments
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Order ID"
// @Security		Bearer
// @Router			/api/payments/order/{id} [get]
// @Success		200	{object}	object
// @Failure		400	{object}	object
// @Failure		500	{object}	object
func (c *PaymentController) ReadByOrderId(ctx *gin.Context) {
	var readIdRequest dto.ReadIdRequest
	if err := ctx.ShouldBindUri(&readIdRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payments, err := c.service.ReadByOrderId(uint(readIdRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, payments)
}
