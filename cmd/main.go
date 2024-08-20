package main

import (
	"api_gateway/api"
	"api_gateway/api/handler"
	"api_gateway/config"
	"api_gateway/logs"
	producer "api_gateway/queue/kafka"
	"api_gateway/service"
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	cfg := config.Load()
	logs.InitLogger()

	fmt.Println(cfg.HTTP_PORT, cfg.AUTH_PORT, cfg.PRODUCT_PORT)
	logs.Logger.Info("Starting the application ...")

	writer := producer.NewKafkaProducer([]string{"kafka:9092"})
	defer writer.Close()

	fmt.Println("not working ")
	srv := service.New()
	h := handler.NewHandler(logs.Logger, srv, writer)
	casbinEnforcer, err := casbin.NewEnforcer("./config/model.conf", "./config/policy.csv")
	if err != nil {
		log.Fatal("Failed to load casbin enforcer", "error", err.Error())
	}
	r := api.Router(h,casbinEnforcer)	
	logs.Logger.Info("Server is running", "PORT", cfg.HTTP_PORT)
	log.Fatal(r.Run(cfg.HTTP_PORT))
}
