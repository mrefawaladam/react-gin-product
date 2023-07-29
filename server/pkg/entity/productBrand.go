package entity

type ProductBrand struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"not null" json:"name"`
	SuplierID uint    `json:"suplier_id"`
	Suplier   Suplier `json:"suplier"`
}
