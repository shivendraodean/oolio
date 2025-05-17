package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	ID int64 `uri:"productId" binding:"required,min=1"`
}

func ValidateGetProductRequest(c *gin.Context) (int64, error) {
	var request ProductRequest

	if err := c.ShouldBindUri(&request); err != nil {
		slog.Error("Invalid product ID supplied", "error", err.Error())
		return 0, err
	}

	return request.ID, nil
}
