package repository

import (
	"github.com/mrefawaladam/server/pkg/entity"
	"gorm.io/gorm"
)

type ProductBrandRepository struct {
	DB *gorm.DB
}

func NewProductBrandRepository(db *gorm.DB) *ProductBrandRepository {
	return &ProductBrandRepository{DB: db}
}

func (r *ProductBrandRepository) GetProductBrandByID(productBrandID uint) (*entity.ProductBrand, error) {
	var productBrand entity.ProductBrand
	if err := r.DB.First(&productBrand, productBrandID).Error; err != nil {
		return nil, err
	}
	return &productBrand, nil
}

func (r *ProductBrandRepository) CreateProductBrand(productBrand *entity.ProductBrand) error {
	return r.DB.Create(productBrand).Error
}
