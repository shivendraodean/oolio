package handler

import (
	"testing"
)

func TestValidateGetProductRequest(t *testing.T) {
	tests := []struct {
		name      string
		productID string
		want      bool
	}{
		{
			name:      "Valid product ID",
			productID: "123",
			want:      true,
		},
		{
			name:      "Empty product ID",
			productID: "",
			want:      false,
		},
		{
			name:      "Whitespace only product ID",
			productID: "   ",
			want:      false,
		},
		{
			name:      "Non-numeric product ID",
			productID: "abc",
			want:      false,
		},
		{
			name:      "Mixed alphanumeric product ID",
			productID: "123abc",
			want:      false,
		},
		{
			name:      "Product ID with special characters",
			productID: "123#456",
			want:      false,
		},
		{
			name:      "Product ID with leading/trailing whitespace",
			productID: "  123  ",
			want:      true,
		},
		{
			name:      "Very large integer product ID",
			productID: "9223372036854775807", // max int64
			want:      true,
		},
		{
			name:      "Too large integer product ID",
			productID: "9223372036854775808", // max int64 + 1
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateGetProductRequest(tt.productID); got != tt.want {
				t.Errorf("ValidateGetProductRequest(%q) = %v, want %v", tt.productID, got, tt.want)
			}
		})
	}
}
