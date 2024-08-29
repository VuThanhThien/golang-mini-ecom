package repositories

import (
	"github.com/VuThanhThien/golang-gorm-postgres/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) GetDB() *gorm.DB {
	return r.db
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) CreateWithTx(tx *gorm.DB, user *models.User) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(user).Error
}

func (r *UserRepository) List(page, pageSize int) ([]models.User, int, error) {
	var users []models.User
	var total int64

	offset := (page - 1) * pageSize

	err := r.db.Model(&models.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Select("users.id").
		Offset(offset).
		Limit(pageSize).
		Find(&users).
		Error
	if err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}
