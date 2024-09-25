package services

import (
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/api/repositories"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models"
	"github.com/VuThanhThien/golang-gorm-postgres/merchant/internal/models/dto"
)

type IInventoryService interface {
	GetInventoryByID(id uint) (*models.Inventory, error)
	GetInventoryByVariantID(variantID uint) (*models.Inventory, error)
	CreateInventory(dto *dto.InventoryDTO) (*models.Inventory, error)
	DeleteInventory(id uint) error
}

type InventoryService struct {
	inventoryRepository *repositories.InventoryRepository
}

func NewInventoryService(inventoryRepository *repositories.InventoryRepository) *InventoryService {
	return &InventoryService{inventoryRepository: inventoryRepository}
}

func (s *InventoryService) GetInventoryByID(id uint) (*models.Inventory, error) {
	return s.inventoryRepository.GetByID(id)
}

func (s *InventoryService) GetInventoryByVariantID(variantID uint) (*models.Inventory, error) {
	return s.inventoryRepository.GetByVariantID(variantID)
}

func (s *InventoryService) CreateInventory(dto *dto.InventoryDTO) (*models.Inventory, error) {
	inventory := &models.Inventory{
		Quantity:  dto.Quantity,
		VariantID: dto.VariantID,
	}
	err := s.inventoryRepository.Create(inventory)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (s *InventoryService) DeleteInventory(id uint) error {
	return s.inventoryRepository.Delete(id)
}
