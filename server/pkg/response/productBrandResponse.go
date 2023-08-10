package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/mrefawaladam/server/pkg/entity"
)

var validate = validator.New()

func init() {
	validate = validator.New()
}

type ProductBrandResponse struct {
	ID        uint             `json:"id"`
	Name      string           `json:"name" validate:"required"`
	SuplierID uint             `json:"suplier_id" validate:"required"`
	Suplier   *SuplierResponse `json:"suplier"`
}

func NewProductBrandResponse(productBrand *entity.ProductBrand) *ProductBrandResponse {
	return &ProductBrandResponse{
		ID:      productBrand.ID,
		Name:    productBrand.Name,
		Suplier: NewSuplierResponse(&productBrand.Suplier), // Menggunakan pointer dari productBrand.Suplier
	}
}
