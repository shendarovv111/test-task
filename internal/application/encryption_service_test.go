package application

import (
	"context"
	"encryption-service/internal/domain/encryption"
	"errors"
	"testing"
	"time"
)

type MemoryCache struct {
	data map[string]string
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]string),
	}
}

func (c *MemoryCache) Get(ctx context.Context, key string) (string, error) {
	if value, ok := c.data[key]; ok {
		return value, nil
	}
	return "", errors.New("cache miss")
}

func (c *MemoryCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	c.data[key] = value
	return nil
}

func TestCachedEncryptionService_Encrypt(t *testing.T) {
	encryptionService := encryption.NewEncryptionService()

	memoryCache := NewMemoryCache()

	service := NewCachedEncryptionService(encryptionService, memoryCache, time.Hour)

	ctx := context.Background()
	memoryCache.Set(ctx, "md5:cached", "cached_result", time.Hour)

	tests := []struct {
		name      string
		ctx       context.Context
		input     string
		algorithm encryption.Algorithm
		want      string
		wantErr   bool
	}{
		{
			name:      "Cache hit",
			ctx:       ctx,
			input:     "cached",
			algorithm: encryption.MD5,
			want:      "cached_result",
			wantErr:   false,
		},
		{
			name:      "Cache miss - MD5",
			ctx:       ctx,
			input:     "test",
			algorithm: encryption.MD5,
			want:      "098f6bcd4621d373cade4e832627b4f6",
			wantErr:   false,
		},
		{
			name:      "Cache miss - SHA256",
			ctx:       ctx,
			input:     "test",
			algorithm: encryption.SHA256,
			want:      "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", // SHA256 хэш для "test"
			wantErr:   false,
		},
		{
			name:      "Unsupported algorithm",
			ctx:       ctx,
			input:     "test",
			algorithm: "unsupported",
			want:      "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Encrypt(tt.ctx, tt.input, tt.algorithm)

			if (err != nil) != tt.wantErr {
				t.Errorf("CachedEncryptionService.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("CachedEncryptionService.Encrypt() = %v, want %v", got, tt.want)
			}

			if !tt.wantErr && tt.name != "Cache hit" {
				cacheKey := string(tt.algorithm) + ":" + tt.input
				cachedValue, err := memoryCache.Get(tt.ctx, cacheKey)
				if err != nil {
					t.Errorf("Value was not cached: %v", err)
				} else if cachedValue != tt.want {
					t.Errorf("Cached value = %v, want %v", cachedValue, tt.want)
				}
			}
		})
	}
}
