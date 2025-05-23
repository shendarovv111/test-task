package encryption

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

type Algorithm string

const (
	MD5 Algorithm = "md5"

	SHA256 Algorithm = "sha256"
)

type Service interface {
	Encrypt(input string, algorithm Algorithm) (string, error)
}

type EncryptionService struct{}

func NewEncryptionService() Service {
	return &EncryptionService{}
}
func (s *EncryptionService) Encrypt(input string, algorithm Algorithm) (string, error) {
	if algorithm == MD5 {
		hash := md5.Sum([]byte(input))
		return hex.EncodeToString(hash[:]), nil
	} else if algorithm == SHA256 {
		hash := sha256.Sum256([]byte(input))
		return hex.EncodeToString(hash[:]), nil
	} else {
		return "", ErrUnsupportedAlgorithm
	}
}
