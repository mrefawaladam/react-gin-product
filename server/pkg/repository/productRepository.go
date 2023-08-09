package repository

import (
	"github.com/mrefawaladam/server/pkg/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetProductByID(productID uint) (*entity.Product, error) {
	var product entity.Product
	if err := r.DB.Preload("ProductBrand").Preload("ProductBrand.Suplier").First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) CreateProduct(product *entity.Product) error {
	return r.DB.Create(product).Error
}
