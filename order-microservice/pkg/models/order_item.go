package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	ID       uint    `gorm:"primaryKey"`
	SellerID uint    `gorm:"type:uint"`
	OrderID  uint    `gorm:"type:int"`
	OfferID  uint    `gorm:"type:int"`
	GTIN     int     `gorm:"type:varchar(20)"`
	Price    float64 `gorm:"type:decimal(10,2)"`
	Quantity int     `gorm:"type:int"`
	SKU      string  `gorm:"type:varchar(255)"`
}

func (OrderItem OrderItem) Migrate(databaseConnection *gorm.DB) error {
	return databaseConnection.AutoMigrate(OrderItem)
}