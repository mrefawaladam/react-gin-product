package usecase

import (
	"github.com/mrefawaladam/server/pkg/entity"
	"github.com/mrefawaladam/server/pkg/repository"
)

type ProductUsecase interface {
	GetProductByID(productID uint) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
}

type productUsecase struct {
	ProductRepo      repository.ProductRepository
	ProductBrandRepo repository.ProductBrandRepository
}

func NewProductUsecase(productRepo repository.ProductRepository, productBrandRepo repository.ProductBrandRepository) ProductUsecase {
	return &productUsecase{
		ProductRepo:      productRepo,
		ProductBrandRepo: productBrandRepo,
	}
}

func (u *productUsecase) GetProductByID(productID uint) (*entity.Product, error) {
	return u.ProductRepo.GetProductByID(productID)
}

func (u *productUsecase) CreateProduct(product *entity.Product) error {
	return u.ProductRepo.CreateProduct(product)
}
