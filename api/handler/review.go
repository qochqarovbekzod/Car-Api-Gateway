package handler

import (
	pb "api_gateway/generated/booking"
	"api_gateway/model"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateReview handles the creation of a new menu item.
// @Summary Create Review
// @Description Create a new Review item
// @Tags Review
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body booking.CreateReviewRequest true "Create Review"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/reviews [post]
func (handler *Handler) CreateReview(ctx *gin.Context) {
	handler.Logger.Info(" creating handler")
	var review pb.CreateReviewRequest
	if err := ctx.ShouldBind(&review); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	req, err := json.Marshal(&review)
	if err != nil {
		handler.Logger.Error("Error marshalling booking request", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err = handler.Producer.ProducerMessage("create-review", req)
	// _, err := handler.Car.CreateReview(ctx, &review)
	if err != nil {
		handler.Logger.Error("Error creating review", err)
		ctx.JSON(500, gin.H{"error": "Error while creating review"})
	}
	handler.Logger.Info("Review created successfully")
	ctx.JSON(201, gin.H{"response": "success response"})
}

// GetReviewById handles the request to fetch a Review item by its ID.
// @Summary Get Review by ID
// @Description Get a Review item by its ID
// @Tags Review
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Review
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/reviews/{id} [get]
func (handler *Handler) GetReviewById(ctx *gin.Context) {
	handler.Logger.Info("GetReviewById called")
	id := ctx.Param("id")

	resp, err := handler.Car.GetByIdReview(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error getting review by id", err)
		ctx.JSON(404, gin.H{"error": "Review not found"})
		return
	}

	handler.Logger.Info("Review found successfully")
	ctx.JSON(200, resp)
}

// UpdateReview handles the update of a menu item.
// @Summary Update Review
// @Description Update an existing Review item
// @Tags Review
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Param Update body booking.UpadateReviewRequest true "Update Review"
// @Success 200 {object} booking.Review
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/reviews/{id} [put]
func (handler *Handler) UpdateReview(ctx *gin.Context) {
	handler.Logger.Info("Updating review...")
	id := ctx.Param("id")

	var review pb.UpadateReviewRequest
	if err := ctx.ShouldBind(&review); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	review.Id = id
	resp, err := handler.Car.UpadateReview(ctx, &review)
	if err != nil {
		handler.Logger.Error("Error updating review", err)
		ctx.JSON(500, gin.H{"error": "Error while updating review"})
		return
	}
	handler.Logger.Info("Review updated successfully")
	ctx.JSON(200, resp)
}

// DeleteReview handles the deletion of a menu item.
// @Summary Delete Review
// @Description Delete an existing review item
// @Tags Review
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} booking.Void
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/reviews/{id} [delete]
func (handler *Handler) DeleteReview(ctx *gin.Context) {
	handler.Logger.Info("Deleting review...")
	id := ctx.Param("id")

	_, err := handler.Car.DeleteReview(ctx, &pb.Id{Id: id})
	if err != nil {
		handler.Logger.Error("Error deleting review", err)
		ctx.JSON(500, gin.H{"error": "Error while deleting review"})
		return
	}
	handler.Logger.Info("Review deleted successfully")
	ctx.JSON(204, gin.H{"success": "Review deleted successfully"})
}

// GetAllReview retrieves a list of review with optional filtering and pagination.
// @Summary Get Review
// @Description Retrieve a list of orders with optional filtering and pagination.
// @Tags Review
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "limit of items to return"
// @Param page query int false "Page for pagination"
// @Param BookingId query string false "Booking identifier to return"
// @Param ProviderId query string false "Provider identifier to return"
// @Param Rating query int false "Rating to return"
// @Param Comment query string false "Comment to return"
// @Success 200 {object} booking.GetAllReviewsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/reviews [get]
func (handler *Handler) GetAllReview(ctx *gin.Context) {
	handler.Logger.Info("GetAllReview called")
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
	var review model.GetAllReviewRequest
	var req pb.GetAllReviewsRequest
	if err := ctx.ShouldBindQuery(&review); err != nil {
		handler.Logger.Error("Error binding request body", err)
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	req.Limit = int32(limit1)
	req.Page = int64(page1)
	req.BookingId = review.BookingId
	req.ProviderId = review.ProviderId
	req.Rating = int64(review.Rating)
	req.Comment = review.Comment
	resp, err := handler.Car.GetAllReviews(ctx, &req)
	if err != nil {
		handler.Logger.Error("Error getting all reviews", err)
		ctx.JSON(500, gin.H{"error": "Error while getting all reviews"})
		return
	}
	handler.Logger.Info("All reviews retrieved successfully")
	ctx.JSON(200, resp)
}
