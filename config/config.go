package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT    string
	AUTH_PORT    string
	PRODUCT_PORT string
	ACCES_TOKEN  string
}

func Load() Config {
	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", "car-api-gateway:8080"))
	cfg.AUTH_PORT = cast.ToString(Coalesce("AUTH_PORT", "auth-servic:50053"))
	cfg.PRODUCT_PORT = cast.ToString(Coalesce("PRODUCT_PORT", "car_sevice:50051"))
	cfg.ACCES_TOKEN = cast.ToString(Coalesce("ACCES_TOKEN", "my_secret_key"))

	return cfg
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
