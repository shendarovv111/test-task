package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"

	"encryption-service/config"
	"encryption-service/internal/application"
	"encryption-service/internal/domain/encryption"
	"encryption-service/internal/infrastructure/cache"
	"encryption-service/internal/interfaces/api"
)

func main() {
	cfg := config.LoadConfig()

	redis := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		DB:   cfg.Redis.DB,
	})

	encryptService := encryption.NewEncryptionService()
	cacheService := cache.NewRedisCache(redis)
	service := application.NewCachedEncryptionService(encryptService, cacheService, cfg.Redis.CacheTTL)

	router := api.SetupRouter(api.NewHandler(service))

	addr := ":" + cfg.Server.Port
	log.Printf("Server starting on port %s", cfg.Server.Port)
	router.Run(addr)
}
