package application

import (
	"context"
	"encryption-service/internal/domain/encryption"
	"encryption-service/internal/infrastructure/cache"
	"fmt"
	"time"
)

type CachedEncryptionService struct {
	encryptionService encryption.Service
	cache             cache.Cache
	cacheTTL          time.Duration
}

func NewCachedEncryptionService(encryptionService encryption.Service, cache cache.Cache, cacheTTL time.Duration) *CachedEncryptionService {
	return &CachedEncryptionService{
		encryptionService: encryptionService,
		cache:             cache,
		cacheTTL:          cacheTTL,
	}
}

func (s *CachedEncryptionService) Encrypt(ctx context.Context, input string, algorithm encryption.Algorithm) (string, error) {
	key := fmt.Sprintf("%s:%s", algorithm, input)

	if result, err := s.cache.Get(ctx, key); err == nil {
		return result, nil
	}

	result, err := s.encryptionService.Encrypt(input, algorithm)
	if err != nil {
		return "", err
	}

	s.cache.Set(ctx, key, result, s.cacheTTL)
	return result, nil
}
