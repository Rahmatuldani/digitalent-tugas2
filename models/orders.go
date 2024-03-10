package models

import (
	"time"

	"gorm.io/gorm"
)

type Items struct {
	Id      uint8  `json:"itemId" gorm:"primaryKey"`
	Code    string `json:"itemCode" gorm:"not null"`
	Desc    string `json:"description" gorm:"not null"`
	Qty     uint8  `json:"quantity" gorm:"not null"`
	OrderId uint8
}

type Orders struct {
	Id        uint8  	`json:"orderId" gorm:"primaryKey"`
	Customer  string 	`json:"customerName" gorm:"not null"`
	OrderedAt time.Time	`json:"orderedAt"`
	Items     []Items	`json:"items" gorm:"foreignKey:OrderId"`
}

type ModelsInterface interface {
	FindAll() ([]Orders, error)
	Create(Orders) error
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

	err := m.Db.Model(&Orders{}).Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (m *ModelsStruct) Create(data Orders) error {
	err := m.Db.Create(&data).Error
	return err
}