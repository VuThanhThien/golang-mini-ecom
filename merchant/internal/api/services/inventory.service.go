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
	DeductQuantity(dto *dto.CreatedOrder) error
	RefundQuantity(dto *dto.CreatedOrder) error
}

type InventoryService struct {
	inventoryRepository *repositories.InventoryRepository
}

func NewInventoryService(inventoryRepository *repositories.InventoryRepository) *InventoryService {
	return &InventoryService{inventoryRepository: inventoryRepository}
}

func (s *InventoryService) GetInventoryByID(id uint) (*models.Inventory, error) {
	var inventory models.Inventory
	err := s.inventoryRepository.GetDB().Where("id = ?", id).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (s *InventoryService) GetInventoryByVariantID(variantID uint) (*models.Inventory, error) {
	return s.inventoryRepository.GetByVariantID(variantID)
}

func (s *InventoryService) CreateInventory(dto *dto.InventoryDTO) (*models.Inventory, error) {
	existingInventory, err := s.inventoryRepository.GetByVariantID(dto.VariantID)
	if err == nil {
		existingInventory.Quantity = dto.Quantity
		err = s.inventoryRepository.Update(existingInventory)
		if err != nil {
			return nil, err
		}
		return existingInventory, nil
	}

	newInventory := &models.Inventory{
		Quantity:  dto.Quantity,
		VariantID: dto.VariantID,
	}
	createdInventory, err := s.inventoryRepository.Create(newInventory)
	if err != nil {
		return nil, err
	}
	return createdInventory, nil
}

func (s *InventoryService) DeleteInventory(id uint) error {
	return s.inventoryRepository.Delete(id)
}

// CreateOrderSucceed deduct product quantity when order was created
func (s *InventoryService) DeductQuantity(dto *dto.CreatedOrder) error {

	for _, item := range dto.Items {
		inventory, err := s.inventoryRepository.GetByVariantID(item.VariantID)
		if err != nil {
			return err
		}
		inventory.Quantity -= item.Quantity
		err = s.inventoryRepository.Update(inventory)
		if err != nil {
			return err
		}
	}

	return nil
}

// RefundOrder refund product was deducted when order was created
func (s *InventoryService) RefundQuantity(dto *dto.CreatedOrder) error {
	for _, item := range dto.Items {
		inventory, err := s.inventoryRepository.GetByVariantID(item.VariantID)
		if err != nil {
			return err
		}
		inventory.Quantity += item.Quantity
		err = s.inventoryRepository.Update(inventory)
		if err != nil {
			return err
		}
	}
	return nil
}
