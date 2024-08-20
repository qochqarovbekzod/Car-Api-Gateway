package api

import (
	"api_gateway/api/handler"
	"api_gateway/api/midleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	_ "api_gateway/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Service
// @version 1.0
// @description API service
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Router(handler *handler.Handler, enfocer *casbin.Enforcer) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(midleware.AuhtMiddleware(enfocer), midleware.Authorization(enfocer))

	user := router.Group("/api/v1/users/profile")
	{
		user.PUT("/:user_id", handler.UpdateUserProfile)
		user.GET("/:user_id", handler.GetUserProfile)
	}

	booking := router.Group("/api/v1/bookings")
	{
		booking.POST("", handler.CreateBooking)
		booking.GET("/:id", handler.GetBookingById)
		booking.PUT("/:id", handler.UpdateBooking)
		booking.DELETE("/:id", handler.DeleteBooking)
		booking.GET("", handler.GetAllBookings)
	}

	provider := router.Group("/api/v1/providers")
	{
		provider.POST("", handler.CreateProvider)
		provider.GET("/:id", handler.GetProviderById)
		provider.PUT("/:id", handler.UpdateProvider)
		provider.DELETE("/:id", handler.DeleteProvider)
		provider.GET("", handler.GetAllProviders)
	}

	service := router.Group("/api/v1/services")
	{
		service.POST("", handler.CreateService)
		service.GET("/:id", handler.GetByIdService)
		service.PUT("/:id", handler.UpdateService)
		service.DELETE("/:id", handler.DeleteService)
		service.GET("", handler.GetAllServices)
	}

	payment := router.Group("/api/v1/payments")
	{
		payment.POST("", handler.CreatePayment)
		payment.GET("/:id", handler.GetPaymentById)
		payment.GET("", handler.GetAllPayment)
	}

	review := router.Group("/api/v1/reviews")
	{
		review.POST("", handler.CreateReview)
		review.GET("/:id", handler.GetReviewById)
		review.PUT("/:id", handler.UpdateReview)
		review.DELETE("/:id", handler.DeleteReview)
		review.GET("", handler.GetAllReview)
	}

	router.GET("/services/best", handler.CareateGet)

	return router

}

