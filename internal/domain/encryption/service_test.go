package encryption
import (
	"testing"
)
func TestEncryptionService_Encrypt(t *testing.T) {
	service := NewEncryptionService()
	tests := []struct {
		name      string
		input     string
		algorithm Algorithm
		want      string
		wantErr   bool
	}{
		{
			name:      "MD5 encryption",
			input:     "test",
			algorithm: MD5,
			want:      "098f6bcd4621d373cade4e832627b4f6",
			wantErr:   false,
		},
		{
			name:      "SHA256 encryption",
			input:     "test",
			algorithm: SHA256,
			want:      "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			wantErr:   false,
		},
		{
			name:      "Unsupported algorithm",
			input:     "test",
			algorithm: "unsupported",
			want:      "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Encrypt(tt.input, tt.algorithm)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptionService.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncryptionService.Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
