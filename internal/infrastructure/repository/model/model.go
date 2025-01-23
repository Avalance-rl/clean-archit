package model

import "time"

type ProductCategory string

var (
	AccessoryEquipment ProductCategory = "accessory_equipment"
	Technical          ProductCategory = "technical"
	RawMaterial        ProductCategory = "raw_material"
)

type Store struct {
	ID          string
	Name        string
	Description string
}

type Product struct {
	ID          string
	Name        string
	StoreID     string
	Price       int
	Description string
	InStock     bool
	Category    ProductCategory
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
