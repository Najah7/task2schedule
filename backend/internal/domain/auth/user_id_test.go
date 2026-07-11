package auth

import (
	"errors"
	"testing"
)

func TestNewUserID(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    UserID
		wantErr error
	}{
		{name: "valid", value: "user-1", want: "user-1"},
		{name: "empty", wantErr: ErrUserIDEmpty},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserID(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewUserID() error = %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("user ID = %q, want %q", got, tt.want)
			}
		})
	}
}
