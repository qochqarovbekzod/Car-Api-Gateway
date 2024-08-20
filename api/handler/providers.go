package handler

import (
	pb "api_gateway/generated/booking"
	"api_gateway/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProvider handles the creation of a new menu item.
// @Summary Create Provider
// @Description Create a new Provider item
// @Tags Provider
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body booking.CreateProvidersRequest true "Create Provider"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/providers [post]
func (handler *Handler) CreateProvider(ctx *gin.Context) {
	handler.Logger.Info("Creating provider")
	var provider pb.CreateProvidersRequest
	if err := ctx.ShouldBind(&provider); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := handler.Car.CreateProviders(ctx, &provider)
	if err != nil {
		fmt.Println(err)
		handler.Logger.Error("Error creating provider", err)
		ctx.JSON(500, gin.H{"error": "Error while creating provider"})
		return
	}

	handler.Logger.Info("Provider created successfully")
	ctx.JSON(201, gin.H{"provider": "provider"})
}

// GetProviderById handles the request to fetch a Provider item by its ID.
// @Summary Get Provider by ID
// @Description Get a Provider item by its ID
// @Tags Provider
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Providers
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/providers/{id} [get]
func (handler *Handler) GetProviderById(ctx *gin.Context) {
	handler.Logger.Info("Getting provider by id")
	id := ctx.Param("id")

	resp, err := handler.Car.GetByIdProviders(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error getting provider by id", err)
		ctx.JSON(404, gin.H{"error": "Provider not found"})
		return
	}

	handler.Logger.Info("Provider found successfully")
	ctx.JSON(200, resp)
}

// UpdateProvider handles the update of a menu item.
// @Summary Update Provider
// @Description Update an existing brovider item
// @Tags Provider
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Param Update body booking.UpdateProvidersRequest true "Update Provider"
// @Success 200 {object} booking.Providers
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/providers/{id} [put]
func (handler *Handler) UpdateProvider(ctx *gin.Context) {
	handler.Logger.Info("Updating provider")
	id := ctx.Param("id")

	var provider pb.UpdateProvidersRequest
	if err := ctx.ShouldBind(&provider); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	provider.Id = id
	resp, err := handler.Car.UpdateProviders(ctx, &provider)
	if err != nil {
		handler.Logger.Error("Error updating provider", err)
		ctx.JSON(500, gin.H{"error": "Error while updating provider"})
		return
	}

	handler.Logger.Info("Provider updated successfully")
	ctx.JSON(200, resp)
}

// DeleteProvider handles the deletion of a menu item.
// @Summary Delete Providers
// @Description Delete an existing providers item
// @Tags Provider
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/providers/{id} [delete]
func (handler *Handler) DeleteProvider(ctx *gin.Context) {
	handler.Logger.Info("Deleting provider")
	id := ctx.Param("id")

	_, err := handler.Car.DeleteProviders(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error deleting provider", err)
		ctx.JSON(500, gin.H{"error": "Error while deleting provider"})
		return
	}

	handler.Logger.Info("Provider deleted successfully")
	ctx.JSON(204, gin.H{" Provider deleted successfully": "deleted successfully"})
}

// GetAllBookings retrieves a list of Ctegory with optional filtering and pagination.
// @Summary Get Providers
// @Description Retrieve a list of orders with optional filtering and pagination.
// @Tags Provider
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "limit of items to return"
// @Param page query int false "Page for pagination"
// @Param UserId query string false "user identifier to return"
// @Param Company_name query string false "company name to return"
// @Param Description query string false "description to return"
// @Param AverageRating query string false "average rating to return"
// @Param Location query string false "location to return from query string"
// @Success 200 {object} booking.GetAllProviderssResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/providers [get]
func (handler *Handler) GetAllProviders(ctx *gin.Context) {
	handler.Logger.Info("Getting all providers")

	limit := ctx.Query("limit")
	page := ctx.Query("page")

	
	var limit1 int
	var page1 int
	var err error
	if limit == ""{
		limit1 = 10
	}
	if page == ""{
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

	var req pb.GetAllProvidersRequest
	var provider model.GetAllProvidersRequest
	if err := ctx.ShouldBindQuery(&provider); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	req.Limit = int32(limit1)
	req.Page = int64(page1)
	req.AverageRating = provider.AverageRating
	req.UserId = provider.UserId
	req.CompanyName = provider.Company_name
	req.Location = provider.Location
	req.Description = provider.Description

	resp, err := handler.Car.GetAllProviderss(ctx, &req)
	if err != nil {
		handler.Logger.Error("Error getting all providers", err)
		ctx.JSON(500, gin.H{"error": "Error while getting all providers"})
		return
	}

	handler.Logger.Info("All providers retrieved successfully")
	ctx.JSON(200, resp)
}
