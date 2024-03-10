package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	Id        uint8  	`json:"orderId" gorm:"primaryKey"`
	Customer  string 	`json:"customerName" gorm:"not null"`
	OrderedAt time.Time	`json:"orderedAt"`
	Items     []Items	`json:"items" gorm:"foreignKey:Id"`
}

type ModelsInterface interface {
	FindAll() ([]Orders, error)
}

type ModelsStruct struct {
	Db *gorm.DB
}

func NewModels(Db *gorm.DB) ModelsInterface {
	Db.AutoMigrate(&Orders{}, &Items{})
	return &ModelsStruct{Db: Db}
}

func (m *ModelsStruct) FindAll() ([]Orders, error) {
	var orders []Orders
	result := []Orders{}

	err := m.Db.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	for _, v := range orders {
		Items := GetItems(v.Id, m.Db)
		result = append(result, Orders{
			Id: v.Id,
			Customer: v.Customer,
			OrderedAt: v.OrderedAt,
			Items: Items,
		})
	}
	return result, nil
}