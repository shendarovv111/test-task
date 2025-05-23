package config

import "os"

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
}

func LoadConfig() *Config {
	return &Config{
		Server: LoadServerConfig(),
		Redis:  LoadRedisConfig(),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
