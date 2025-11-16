package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJson(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return false
	}
	return true
}
