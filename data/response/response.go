package response

import "time"

type ItemResponse struct {
	Id      uint8  `json:"itemId"`
	Code    string `json:"itemCode"`
	Desc    string `json:"description"`
	Qty     uint8  `json:"quantity"`
}

type OrderResponse struct {
	Id        uint8          `json:"orderId"`
	Customer  string         `json:"customerName"`
	OrderedAt time.Time      `json:"orderedAt"`
	Items     []ItemResponse `json:"items"`
}

type WebResponse struct {
	Status	string 		`json:"status"`
	Message	string 		`json:"message"`
	Data	interface{} `json:"data,omitempty"`
}