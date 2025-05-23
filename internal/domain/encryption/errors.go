package encryption

import (
	"errors"
	"fmt"
)

var ErrUnsupportedAlgorithm = errors.New("unsupported encryption algorithm")
var ErrEmptyInput = errors.New("input string cannot be empty")

// WrapError оборачивает ошибку с дополнительным контекстом
func WrapError(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %w", msg, err)
}

// NewError создает новую ошибку с форматированным сообщением
func NewError(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
