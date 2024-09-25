package repositories

import (
	"gorm.io/gorm"
)

// BaseRepository interface
type BaseRepository[T any] interface {
	Create(entity *T) error
	GetByID(id uint) (*T, error)
	Update(entity *T) error
	Delete(id uint) error
	GetDB() *gorm.DB
}

// BaseRepositoryImpl struct
type BaseRepositoryImpl[T any] struct {
	db *gorm.DB
}

// NewBaseRepository creates a new base repository
func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &BaseRepositoryImpl[T]{db: db}
}

// Implement BaseRepository methods
func (r *BaseRepositoryImpl[T]) GetDB() *gorm.DB {
	return r.db
}

func (r *BaseRepositoryImpl[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepositoryImpl[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return &entity, err
}

func (r *BaseRepositoryImpl[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepositoryImpl[T]) Delete(id uint) error {
	return r.db.Delete(new(T), id).Error
}
