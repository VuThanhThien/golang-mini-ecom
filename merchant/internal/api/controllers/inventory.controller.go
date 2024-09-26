package controllers

import (
	"net/http"
	"strconv"

	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/services"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	inventoryService services.IInventoryService
}

func NewInventoryController(inventoryService services.IInventoryService) *InventoryController {
	return &InventoryController{inventoryService: inventoryService}
}

// CreateInventory godoc
// @Summary Create a new inventory
// @Description Create a new inventory with the given details
// @Tags Inventory
// @Accept json
// @Produce json
// @Param inventory body dto.InventoryDTO true "Inventory details"
// @Success 201 {object} dto.InventoryDTO
// @Router /inventory [post]
func (c *InventoryController) CreateInventory(ctx *gin.Context) {
	var inventoryDTO dto.InventoryDTO
	if err := ctx.ShouldBindJSON(&inventoryDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inventory, err := c.inventoryService.CreateInventory(&inventoryDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, inventory)
}

// GetInventoryByID godoc
// @Summary Get an inventory by ID
// @Description Get an inventory by its ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param id path string true "Inventory ID"
// @Success 200 {object} dto.InventoryDTO
// @Router /inventory/{id} [get]
func (c *InventoryController) GetInventoryByID(ctx *gin.Context) {
	var readIdRequest dto.ReadIdRequest
	if err := ctx.ShouldBindUri(&readIdRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	inventory, err := c.inventoryService.GetInventoryByID(uint(readIdRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, inventory)
}

// GetInventoryByVariantID godoc
// @Summary Get an inventory by Variant ID
// @Description Get an inventory by its Variant ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param id path string true "Variant ID"
// @Success 200 {object} dto.InventoryDTO
// @Router /inventory/variant/{id} [get]
func (c *InventoryController) GetInventoryByVariantID(ctx *gin.Context) {
	var readIdRequest dto.ReadIdRequest
	if err := ctx.ShouldBindUri(&readIdRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	inventory, err := c.inventoryService.GetInventoryByVariantID(uint(readIdRequest.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, inventory)
}

// DeleteInventory godoc
// @Summary Delete an inventory by ID
// @Description Delete an inventory by its ID
// @Tags Inventory
// @Accept json
// @Produce json
// @Param id path string true "Inventory ID"
// @Success 204
// @Router /inventory/{id} [delete]
func (c *InventoryController) DeleteInventory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = c.inventoryService.DeleteInventory(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
