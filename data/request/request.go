package request

import "time"

type CreateItemRequest struct {
	Code string `json:"itemCode" validate:"required"`
	Desc string `json:"description" validate:"required"`
	Qty  uint8  `json:"quantity" validate:"required"`
}

type CreateOrderRequest struct {
	Customer string              `json:"customerName" validate:"required"`
	Items    []CreateItemRequest `json:"items" validate:"required"`
}

type UpdateItemRequest struct {
	Id 	uint8 	`json:"lineItemId" validate:"required"`
	Code string `json:"itemCode" validate:"required"`
	Desc string `json:"description" validate:"required"`
	Qty  uint8  `json:"quantity" validate:"required"`
}

type UpdateOrderRequest struct {
	Customer  string              `json:"customerName" validate:"required"`
	OrderedAt time.Time        	  `json:"orderedAt"`
	Items     []UpdateItemRequest `json:"items" validate:"required"`
}