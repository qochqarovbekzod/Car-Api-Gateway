package model

type GetAllBookingRequest struct {
	UserId string `json:"user_id"`
	ProviderId string	`json:"provider_id"`
	ServiceId string  `json:"service_id"`
	Status string  `json:"status"`
	TotalPrice float32 `json:"total_price"`

}

type GetAllServicesRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Duration int64 `json:"duration"`
}

type GetAllReviewRequest struct {
	BookingId string `json:"booking_id"`
	ProviderId string `json:"provider_id"`
	Rating int `json:"rating"`
	Comment string `json:"comment"`
}

type GetAllProvidersRequest struct {
	UserId string `json:"user_id"`
	Company_name string `json:"company_name"`
	Description string `json:"description"`
	AverageRating float32 `json:"average_rating"`
	Location string `json:"location"`
}

