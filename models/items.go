package models

import "gorm.io/gorm"



func GetItems(id uint8, Db *gorm.DB) []Items {
	items := []Items{}

	Db.Where("order_id = ?", id).Find(&items)

	return items
}

// func CreateItem(item Items, Db *gorm.DB) error {
// 	err := Db.Create(&item)
// 	return nil
// }