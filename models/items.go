package models

import "gorm.io/gorm"

type Items struct {
	Id      uint8  `json:"itemId" gorm:"primaryKey"`
	Code    string `json:"itemCode" gorm:"not null"`
	Desc    string `json:"description" gorm:"not null"`
	Qty     uint8  `json:"quantity" gorm:"not null"`
	OrderId uint8
}

func GetItems(id uint8, Db *gorm.DB) []Items {
	items := []Items{}

	Db.Where("order_id = ?", id).Find(&items)

	return items
}