package response

import (
	"github.com/mrefawaladam/server/pkg/entity"
)

type ProductBrandResponse struct {
	ID        uint             `json:"id"`
	Name      string           `json:"name"`
	SuplierID uint             `json:"suplier_id"`
	Suplier   *SuplierResponse `json:"suplier"`
}

func NewProductBrandResponse(productBrand *entity.ProductBrand) *ProductBrandResponse {
	return &ProductBrandResponse{
		ID:      productBrand.ID,
		Name:    productBrand.Name,
		Suplier: NewSuplierResponse(&productBrand.Suplier),
	}
}
