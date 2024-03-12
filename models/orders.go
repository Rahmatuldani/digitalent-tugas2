package models

import (
	"time"

	"github.com/Rahmatuldani/assignment2/data/request"
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
	Update(uint8, request.UpdateOrderRequest) (Orders, error)
	Delete(uint8) (int, string, error)
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

func (m *ModelsStruct) Update(id uint8, data request.UpdateOrderRequest) (Orders, error) {
	var order Orders
	if err := m.Db.First(&order, id).Error; err != nil {
		return order, err
	}
	
	for _, v := range data.Items{
		if err := m.Db.Where("order_id = ? AND id = ?", id, v.Id).Save(&Items{Code: v.Code, Desc: v.Desc, Qty: v.Qty,}).Error; err != nil {
			return order, err
		}
	}
	if err := m.Db.Where("id = ?", id).Save(&Orders{Customer: data.Customer, OrderedAt: data.OrderedAt}).Error; err != nil {
		return order, err
	}
	
	return order, nil
}

func (m *ModelsStruct) Create(data Orders) error {
	err := m.Db.Create(&data).Error
	return err
}

func (m *ModelsStruct) Delete(id uint8) (int, string, error) {
	var order Orders
	if err := m.Db.First(&order, id).Error; err != nil {
		return 404, "Order not found", err
	}

	if err := m.Db.Where("order_id = ?", id).Delete(&Items{}).Error; err != nil {
		return 500, "Deleted Failed", err
	}
	if err := m.Db.Delete(&Orders{}, id).Error; err != nil {
		return 500, "Deleted Failed", err
	}
	return 200, "Deleted success", nil
}