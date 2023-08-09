package entity

type Product struct {
	ID             uint         `gorm:"primaryKey" json:"id"`
	Name           string       `gorm:"not null" json:"name"`
	ProductBrandID uint         `json:"product_brand_id"`
	ProductBrand   ProductBrand `json:"product_brand"`
}
