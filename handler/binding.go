package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const errInvalidJSON = "잘못된 요청입니다"

func BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidJSON})
		return false
	}
	return true
}
