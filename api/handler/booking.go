package handler

import (
	pb "api_gateway/generated/booking"
	"api_gateway/model"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBooking handles the creation of a new menu item.
// @Summary Create Booking
// @Description Create a new Booking item
// @Tags Booking
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body booking.CreateBookingRequest true "Create Booking"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/bookings [post]
func (handler *Handler) CreateBooking(ctx *gin.Context) {
	handler.Logger.Info("Creating booking...")
	var booking pb.CreateBookingRequest
	if err := ctx.ShouldBind(&booking); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	req,err:=json.Marshal(&booking)
	if err != nil {
		handler.Logger.Error("Error marshalling booking request", err)
        ctx.JSON(400, gin.H{"error": "Invalid request body"})
        return
    }
	err=handler.Producer.ProducerMessage("create-booking",req)

	// _, err := handler.Car.CreateBooking("create-booking", &booking)
	if err != nil {
		fmt.Println(err)
		handler.Logger.Error("Error creating booking", err)
		ctx.JSON(500, gin.H{"error": "Error while creating booking"})
		return
	}

	handler.Logger.Info("Booking created successfully")
	ctx.JSON(201, gin.H{"booking": "created successfully"})
}

// GetBookingById handles the request to fetch a category item by its ID.
// @Summary Get Booking by ID
// @Description Get a Booking item by its ID
// @Tags Booking
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Booking
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/bookings/{id} [get]
func (handler *Handler) GetBookingById(ctx *gin.Context) {
	handler.Logger.Info("Getting booking by id...")
	id := ctx.Param("id")

	resp, err := handler.Car.GetByIdBooking(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error getting booking by id", err)
		ctx.JSON(404, gin.H{"error": "Booking not found"})
		return
	}

	handler.Logger.Info("Booking found successfully")
	fmt.Println(resp)
	ctx.JSON(200, resp)
}

// UpdateBooking handles the update of a menu item.
// @Summary Update Booking
// @Description Update an existing booking item
// @Tags Booking
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Param Update body booking.UpdateBookingRequest true "Update Booking"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/bookings/{id} [put]
func (handler *Handler) UpdateBooking(ctx *gin.Context) {
	handler.Logger.Info("Updating booking...")
	id := ctx.Param("id")

	var booking pb.UpdateBookingRequest
	if err := ctx.ShouldBind(&booking); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	booking.Id = id

	req,err:=json.Marshal(&booking)
	if err != nil {
		handler.Logger.Error("Error marshalling booking request", err)
        ctx.JSON(400, gin.H{"error": "Invalid request body"})
        return
    }
	err=handler.Producer.ProducerMessage("update-booking",req)
	// resp, err := handler.Car.UpdateBooking(ctx, &booking)
	if err != nil {
		handler.Logger.Error("Error updating booking", err)
		ctx.JSON(500, gin.H{"error": "Error while updating booking"})
		return
	}

	handler.Logger.Info("Booking updated successfully")
	ctx.JSON(200, gin.H{"success": "Booking updated successfully"})
}

// DeleteBooking handles the deletion of a menu item.
// @Summary Delete Booking
// @Description Delete an existing booking item
// @Tags Booking
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/bookings/{id} [delete]
func (handler *Handler) DeleteBooking(ctx *gin.Context) {
	handler.Logger.Info("Deleting booking...")
	id := ctx.Param("id")
	req,err:=json.Marshal(&id)
	if err != nil {
		handler.Logger.Error("Error marshalling booking request", err)
        ctx.JSON(400, gin.H{"error": "Invalid request body"})
        return
    }
	err=handler.Producer.ProducerMessage("delete-boking",req)
	// _, err := handler.Car.DeleteBooking(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error deleting booking", err)
		ctx.JSON(500, gin.H{"error": "Error while deleting booking"})
		return
	}

	handler.Logger.Info("Booking deleted successfully")
	ctx.JSON(204, gin.H{"booking": "created successfully"})
}

// GetAllBookings retrieves a list of Ctegory with optional filtering and pagination.
// @Summary Get Booking
// @Description Retrieve a list of orders with optional filtering and pagination.
// @Tags Booking
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "limit of items to return"
// @Param page query int false "Page for pagination"
// @Param UserId query string false "user identifier to return"
// @Param ProviderId query string false "provider identifier to return"
// @Param ServiceId query string false "service identifier to return"
// @Param Status query string false "status to return" (enum :pending, confirmed, completed, cancelled)
// @Param TotalPrice query string false "total price to return"
// @Success 200 {object} booking.GetAllBookingsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/bookings [get]
func (handler *Handler) GetAllBookings(ctx *gin.Context) {
	fmt.Println("salomsdfss dfgds")
	handler.Logger.Info("Getting all bookings...")
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	var limit1 int
	var page1 int
	var err error

	if limit != "" {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			handler.Logger.Error("Error parsing limit query parameter", err)
			ctx.JSON(400, gin.H{"error": "Invalid limit query parameter"})
			return
		}

	}
	if page != "" {
		page1, err = strconv.Atoi(page)
		if err != nil {
			handler.Logger.Error("Error parsing page query parameter", err)
			ctx.JSON(400, gin.H{"error": "Invalid page query parameter"})
			return
		}
	}

	var booking model.GetAllBookingRequest
	var req pb.GetAllBookingRequest
	if err := ctx.ShouldBindQuery(&booking); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	req.Limit = int64(limit1)
	req.Page = int64(page1)
	req.UserId = booking.UserId
	req.ProviderId = booking.ProviderId
	req.ServiceId = booking.ServiceId
	req.Status = booking.Status
	req.TotalPrice = booking.TotalPrice

	resp, err := handler.Car.GetAllBookings(ctx, &req)
	fmt.Println("defovnekdwjls;adfg")
	if err != nil {
		fmt.Println(err)
		handler.Logger.Error("Error getting all bookings", err)
		ctx.JSON(500, gin.H{"error": "Error while getting all bookings"})
		return
	}
	handler.Logger.Info("All bookings retrieved successfully")
	ctx.JSON(200, resp)
}
