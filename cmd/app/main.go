package main

import (
	"log"

	"github.com/ccxnu/ips-redis-mysql/config"
	"github.com/ccxnu/ips-redis-mysql/internal/service"
)

func main() {
	// Load configuration
	config.InitConfig()

	// Start the server
	redisService := service.NewRedisService(config.AppConfig)
	err := redisService.Run()
	if err != nil {
		log.Fatalf("Error running Redis service: %v", err)
	}

}
