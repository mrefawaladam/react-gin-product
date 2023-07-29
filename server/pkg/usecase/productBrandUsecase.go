package usecase

import (
	"github.com/mrefawaladam/server/pkg/entity"
	"github.com/mrefawaladam/server/pkg/repository"
)

type ProductBrandUsecase interface {
	GetProductBrandByID(productBrandID uint) (*entity.ProductBrand, error)
	CreateProductBrand(productBrand *entity.ProductBrand) error
}

type productBrandUsecase struct {
	ProductBrandRepo repository.ProductBrandRepository
}

func NewProductBrandUsecase(productBrandRepo repository.ProductBrandRepository) ProductBrandUsecase {
	return &productBrandUsecase{ProductBrandRepo: productBrandRepo}
}

func (u *productBrandUsecase) GetProductBrandByID(productBrandID uint) (*entity.ProductBrand, error) {
	return u.ProductBrandRepo.GetProductBrandByID(productBrandID)
}

func (u *productBrandUsecase) CreateProductBrand(productBrand *entity.ProductBrand) error {
	return u.ProductBrandRepo.CreateProductBrand(productBrand)
}
