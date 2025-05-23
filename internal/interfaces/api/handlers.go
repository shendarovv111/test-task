package api

import (
	"context"
	"encryption-service/internal/domain/encryption"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EncryptionService interface {
	Encrypt(ctx context.Context, input string, algorithm encryption.Algorithm) (string, error)
}

type Handler struct {
	service EncryptionService
}

func NewHandler(service EncryptionService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) EncryptHandler(c *gin.Context) {
	var req EncryptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: fmt.Sprintf("Неверный формат запроса: %s", err.Error()),
		})
		return
	}

	if req.Input == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Входная строка не может быть пустой",
		})
		return
	}

	var algorithm encryption.Algorithm
	switch req.Algorithm {
	case "md5":
		algorithm = encryption.MD5
	case "sha256":
		algorithm = encryption.SHA256
	default:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: fmt.Sprintf("Неподдерживаемый алгоритм: %s. Используйте 'md5' или 'sha256'", req.Algorithm),
		})
		return
	}

	result, err := h.service.Encrypt(c, req.Input, algorithm)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, encryption.ErrUnsupportedAlgorithm) {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, EncryptResponse{Result: result})
}
