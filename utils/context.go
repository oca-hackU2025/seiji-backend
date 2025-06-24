package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromContext(c *gin.Context) (uint, error) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("userID not found in context")
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		return 0, errors.New("userID in context is not uint")
	}

	return userID, nil
}
