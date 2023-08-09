package response

import (
	"github.com/mrefawaladam/server/pkg/entity"
)

type ProductResponse struct {
	ID             uint                  `json:"id"`
	Name           string                `json:"name"`
	ProductBrandID uint                  `json:"product_brand_id"`
	ProductBrand   *ProductBrandResponse `json:"product_brand"`
}

func NewProductResponse(product *entity.Product) *ProductResponse {
	return &ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		ProductBrandID: product.ProductBrandID,
		ProductBrand: &ProductBrandResponse{
			ID:   product.ProductBrand.ID,
			Name: product.ProductBrand.Name,
		},
	}
}
