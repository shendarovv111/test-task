package config

import (
	"strconv"
	"time"
)

type RedisConfig struct {
	Host     string
	Port     string
	DB       int
	CacheTTL time.Duration
}

func LoadRedisConfig() RedisConfig {
	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	cacheTTL, _ := strconv.Atoi(getEnv("CACHE_TTL", "3600"))

	return RedisConfig{
		Host:     getEnv("REDIS_HOST", "localhost"),
		Port:     getEnv("REDIS_PORT", "6379"),
		DB:       redisDB,
		CacheTTL: time.Duration(cacheTTL) * time.Second,
	}
}
