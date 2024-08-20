package handler

import (
	"log/slog"

	"api_gateway/generated/auth"
	"api_gateway/generated/booking"
	"api_gateway/service"

	"api_gateway/queue/kafka"
)

type Handler struct {
	Logger *slog.Logger
	Auth   auth.AuthServiceClient
	Car  booking.BookingServiceClient
	Producer producer.KafkaProducer
}

func NewHandler(logger *slog.Logger, srv service.ServiceMnager, prod producer.KafkaProducer) *Handler {
	
	return &Handler{
        Logger: logger,
        Auth:   srv.User(),
		Car: srv.Booking(),
		Producer: prod,
    }
}
