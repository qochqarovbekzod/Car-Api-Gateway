package handler

import (
	pb "api_gateway/generated/booking"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatePayment handles the creation of a new menu item.
// @Summary Create Payment
// @Description Create a new Payment item
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body booking.CreatePaymentRequest true "Create Payment"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/payments [post]
func (handler *Handler) CreatePayment(ctx *gin.Context) {
	handler.Logger.Info("CreatePayment called with payload")
	var payment pb.CreatePaymentRequest
	if err := ctx.ShouldBind(&payment); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := handler.Car.CreatePayment(ctx, &payment)
	if err != nil {
		handler.Logger.Error("Error creating payment", err)
		ctx.JSON(500, gin.H{"error": "Error while creating payment"})
		return
	}
	handler.Logger.Info("Payment created successfully")
	ctx.JSON(201, gin.H{"success": "Payment created successfully"})
}

// GetPaymentById handles the request to fetch a category item by its ID.
// @Summary Get Payment by ID
// @Description Get a Payment item by its ID
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Payment
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/payments/{id} [get]
func (handler *Handler) GetPaymentById(ctx *gin.Context) {
	handler.Logger.Info("GetPaymentById")
	id := ctx.Param("id")

	resp, err := handler.Car.GetByIdPayment(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error getting payment by id", err)
		ctx.JSON(404, gin.H{"error": "Payment not found"})
		return
	}

	handler.Logger.Info("Payment found successfully")
	ctx.JSON(200, resp)
}

// GetAllPayment retrieves a list of Ctegory with optional filtering and pagination.
// @Summary Get Payment
// @Description Retrieve a list of orders with optional filtering and pagination.
// @Tags Payment
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} booking.GetAllPaymentsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/payments [get]
func (handler *Handler) GetAllPayment(ctx *gin.Context) {
	handler.Logger.Info("GetAllPayment called")

	resp, err := handler.Car.GetAllPayments(ctx, &pb.Void{})
	if err != nil {
		handler.Logger.Error("Error getting all payments", err)
		ctx.JSON(500, gin.H{"error": "Error while getting all payments"})
		return
	}
	handler.Logger.Info("All payments retrieved successfully")
	ctx.JSON(200, resp)
}

// CareateGet retrieves a list of Ctegory with optional filtering and pagination.
// @Summary Get Services
// @Description Retrieve a list of orders with optional filtering and pagination.
// @Tags Services
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} booking.Service
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /services/best [get]
func (handler *Handler) CareateGet(ctx *gin.Context) {
	handler.Logger.Info("CareateGet called")
	id := uuid.NewString()
	ctx.JSON(200, pb.Service{
		Id:          id,
		Name:        "Best Service",
		Description: "Best Service Description",
		Price:       100,
		CreatedAt:   "2022/01/01",
		UpdatedAt:   "2022/01/01"})
	if true {

		return
	}

	resp, err := handler.Car.CreateGet(ctx, &pb.Void{})
	if err != nil {
		fmt.Println(err)
		handler.Logger.Error("Error getting", err)
		ctx.JSON(500, gin.H{"error": "Error while getting"})
		return
	}
	handler.Logger.Info("Get successfully")
	ctx.JSON(200, resp)
}
