package handler

import (
	"strconv"
	"strings"
)

func ValidateGetProductRequest(productID string) bool {
	if productID == "" {
		return false
	}

	trimmedID := strings.TrimSpace(productID)
	if trimmedID == "" {
		return false
	}

	_, err := strconv.ParseInt(trimmedID, 10, 64)
	if err != nil {
		return false
	}

	return true
}
