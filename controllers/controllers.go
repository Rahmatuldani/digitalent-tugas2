package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Rahmatuldani/assignment2/data/request"
	"github.com/Rahmatuldani/assignment2/data/response"
	"github.com/Rahmatuldani/assignment2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ControllerStruct struct {
	model models.ModelsInterface
	validate *validator.Validate
}

func NewController(model models.ModelsInterface, validate *validator.Validate) *ControllerStruct{
	return &ControllerStruct{
		model: model,
		validate: validate,
	}
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
	fmt.Println("Here1")
	result, err := m.model.FindAll()
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Error: err.Error()})
		return
	}
	
	ctx.JSON(200, result)
}

// Orders godoc
// @Summary Create new order
// @Schemes
// @Description Create new orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param req body request.CreateOrderRequest true "Request Body"
// @Success 200 {object} models.Orders
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /orders [post]
func (m *ControllerStruct) CreateOrder(ctx *gin.Context) {
	var req models.Orders
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{Error: err.Error()})
		return
	}
	if err := m.validate.Struct(req); err != nil {
		ctx.JSON(400, response.ErrorResponse{Error: err.Error()})
		return	
	}
	
	req.OrderedAt = time.Now()
	err := m.model.Create(req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Error: err.Error()})
		return
	}
	
	ctx.JSON(200, req)
}

// Orders godoc
// @Summary Update order
// @Schemes
// @Description Update orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Id"
// @Param req body request.UpdateOrderRequest true "Request Body"
// @Success 200 {object} models.Orders
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /orders/{id} [put]
func (m *ControllerStruct) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Status: "Bad Request",
			Message: "Param id required",
			Error: err.Error()})
		return
	}

	var req request.UpdateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{Error: err.Error()})
		return
	}
	if err := m.validate.Struct(req); err != nil {
		ctx.JSON(400, response.ErrorResponse{Error: err.Error()})
		return	
	}
	
	order, err := m.model.Update(uint8(aid), req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{Error: err.Error()})
		return
	}
	
	ctx.JSON(200, order)
}

// Orders godoc
// @Summary 	Delete order
// @Schemes
// @Description Delete order
// @Tags 		Orders
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "Id"
// @Success 	200 {object} response.WebResponse
// @Failure 	400 {object} response.ErrorResponse
// @Failure 	500 {object} response.ErrorResponse
// @Router 		/orders/{id} [delete]
func (m *ControllerStruct) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Status: "Bad Request",
			Message: "Param id required",
			Error: err.Error()})
		return
	}
	
	code, message, err := m.model.Delete(uint8(aid))
	if err != nil {
		ctx.JSON(code, response.ErrorResponse{
			Status: "Error",
			Message: message,
			Error: err.Error()})
		return
	}
	
	ctx.JSON(code, response.WebResponse{
		Status: "OK",
		Message: message,
	})
}