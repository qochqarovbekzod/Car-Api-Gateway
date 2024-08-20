package service

import (
	"api_gateway/config"
	"api_gateway/generated/auth"
	"api_gateway/generated/booking"
	"api_gateway/logs"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceMnager interface {
	User() auth.AuthServiceClient
	Booking() booking.BookingServiceClient
}

type serviceManagerImpl struct {
	authClient    auth.AuthServiceClient
	bookingClient booking.BookingServiceClient
}

func New() ServiceMnager {
	cfg := config.Load()
	logs.InitLogger()
	fmt.Printf("/n/n/n/n/n%s %s %s/n/n/n/n",cfg.AUTH_PORT,cfg.PRODUCT_PORT,cfg.HTTP_PORT)

	connAuth, err := grpc.NewClient(cfg.AUTH_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.Logger.Error("Failed to created Auth client connection", "error", err.Error())
		log.Println(err)
		return nil
	}

	authSercice := auth.NewAuthServiceClient(connAuth)

	connBooking, err := grpc.NewClient(cfg.PRODUCT_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.Logger.Error("Failed to created Booking client connection", "error", err.Error())
		log.Println(err)
		return nil
	}
	bookingService := booking.NewBookingServiceClient(connBooking)

	return &serviceManagerImpl{
		authClient: authSercice,
		bookingClient: bookingService,
	}
}

func (s serviceManagerImpl) User() auth.AuthServiceClient {
	return s.authClient
}

func (s serviceManagerImpl) Booking() booking.BookingServiceClient {
	return s.bookingClient
}