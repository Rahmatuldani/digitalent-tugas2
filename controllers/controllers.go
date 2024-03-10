package controllers

import (
	"github.com/Rahmatuldani/assignment2/data/response"
	"github.com/Rahmatuldani/assignment2/models"
	"github.com/gin-gonic/gin"
)

type ControllerStruct struct {
	model models.ModelsInterface
}

func NewController(model models.ModelsInterface) *ControllerStruct{
	return &ControllerStruct{model: model}
}

// Orders godoc
// @Summary Get all orders
// @Schemes
// @Description Get all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} []models.Orders
// @Failure 500 {object} response.ErrorResponse
// @Router /orders [get]
func (m *ControllerStruct) GetOrders(ctx *gin.Context) {
	result, err := m.model.FindAll()
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Error: err.Error()})
		return
	}
	
	ctx.JSON(200, result)
}