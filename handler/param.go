package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const errInvalidID = "잘못된 유저 ID입니다"

func ParseUserID(c *gin.Context) (uint, bool) {
	userIDStr := c.Param("user_id")
	var userID uint
	_, err := fmt.Sscanf(userIDStr, "%d", &userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidID})
		return 0, false
	}
	return userID, true
}
