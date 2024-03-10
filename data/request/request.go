package request

type CreateItemRequest struct {
	Code    string `json:"itemCode" validate:"required"`
	Desc    string `json:"description" validate:"required"`
	Qty     uint8  `json:"quantity" validate:"required"`
}

type CreateOrderRequest struct {
	Customer  string 				`json:"customerName" validate:"required"`
	Items     []CreateItemRequest	`json:"items" validate:"required"`
}