package handler

import (
	pb "api_gateway/generated/booking"
	"api_gateway/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateService handles the creation of a new menu item.
// @Summary Create Service
// @Description Create a new Service item
// @Tags Service
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body booking.CreateServiceRequest true "Create Service"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/services [post]
func (handler *Handler) CreateService(ctx *gin.Context) {
	handler.Logger.Info("Creating service")
	var service pb.CreateServiceRequest
	if err := ctx.ShouldBind(&service); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := handler.Car.CreateService(ctx, &service)
	if err != nil {
		handler.Logger.Error("Error creating service", err)
		ctx.JSON(500, gin.H{"error": "Error while creating service"})
		return
	}

	handler.Logger.Info("Service created successfully")
	ctx.JSON(201, gin.H{"service": "Service created successfully"})
}

// GetByIdService handles the request to fetch a Service item by its ID.
// @Summary Get Service by ID
// @Description Get a Service item by its ID
// @Tags Service
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Service
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/services/{id} [get]
func (handler *Handler) GetByIdService(ctx *gin.Context) {
	handler.Logger.Info("Getting service by id")
	id := ctx.Param("id")

	resp, err := handler.Car.GetByIdService(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error getting service by id", err)
		ctx.JSON(404, gin.H{"error": "Service not found"})
		return
	}

	handler.Logger.Info("Service found successfully")
	ctx.JSON(200, resp)
}

// UpdateService handles the update of a menu item.
// @Summary Update Service
// @Description Update an existing Service item
// @Tags Service
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Param Update body booking.UpdateServiceRequest true "Update Service"
// @Success 200 {object} booking.Service
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/services/{id} [put]
func (handler *Handler) UpdateService(ctx *gin.Context) {
	handler.Logger.Info("Updating service")
	id := ctx.Param("id")

	var service pb.UpdateServiceRequest
	if err := ctx.ShouldBind(&service); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	service.Id = id
	resp, err := handler.Car.UpdateService(ctx, &service)
	if err != nil {
		handler.Logger.Error("Error updating service", err)
		ctx.JSON(500, gin.H{"error": "Error while updating service"})
		return
	}

	handler.Logger.Info("Service updated successfully")
	ctx.JSON(200, resp)
}

// DeleteService handles the deletion of a menu item.
// @Summary Delete Service
// @Description Delete an existing service item
// @Tags Service
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/services/{id} [delete]
func (handler *Handler) DeleteService(ctx *gin.Context) {
	handler.Logger.Info("Deleting service")
	id := ctx.Param("id")

	_, err := handler.Car.DeleteService(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error deleting service", err)
		ctx.JSON(500, gin.H{"error": "Error while deleting service"})
		return
	}

	handler.Logger.Info("Service deleted successfully")
	ctx.JSON(204, gin.H{"success": "Service deleted successfully"})
}

// GetAllServices retrieves a list of Ctegory with optional filtering and pagination.
// @Summary Get Service
// @Description Retrieve a list of orders with optional filtering and pagination.
// @Tags Service
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "limit of items to return"
// @Param page query int false "Page for pagination"
// @Param Name query string false "user identifier to return"
// @Param Description query string false "provider identifier to return"
// @Param Price query float64 false "service identifier to return"
// @Param Duration query int64 false "total price to return"
// @Success 200 {object} booking.GetAllServicesResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/services [get]
func (handler *Handler) GetAllServices(ctx *gin.Context) {
	handler.Logger.Info("Getting all services...")
	limit := ctx.Query("limit")
	page := ctx.Query("page")

	var limit1 int
	var page1 int
	var err error
	if limit == "" || page == "" {
		limit1 = 10
		page1 = 1
	}
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
	var service model.GetAllServicesRequest
	var req pb.GetAllServicesRequest
	service.Name = ctx.Query("Name")
	service.Description = ctx.Query("Description")
	price := ctx.Query("Price")
	if price != "" {
		service.Price, err = strconv.ParseFloat(price, 64)
		if err != nil {
			handler.Logger.Error("Error parsing page query parameter", err)
			ctx.JSON(400, gin.H{"error": "Invalid page query parameter"})
			return
		}
	}

	duration := ctx.Query("Duration")
	if duration != "" {
		service.Duration, err = strconv.ParseInt(duration, 10, 64)
		if err != nil {
			handler.Logger.Error("Error parsing page query parameter", err)
			ctx.JSON(400, gin.H{"error": "Invalid page query parameter"})
			return
		}
	}

	req.Limit = int32(limit1)
	req.Page = int64(page1)
	req.Name = service.Name
	req.Description = service.Description
	req.Price = float32(service.Price)
	req.Duration = service.Duration

	resp, err := handler.Car.GetAllServices(ctx, &req)
	if err != nil {
		handler.Logger.Error("Error getting all services", err)
		ctx.JSON(500, gin.H{"error": "Error while getting all services"})
		return
	}

	handler.Logger.Info("All services found successfully")
	ctx.JSON(200, resp)
}
