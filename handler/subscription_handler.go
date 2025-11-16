package handler

import (
	"fmt"
	"net/http"

	"github.com/chaeyoungeee/blog-feed-notifier/dto"
	"github.com/chaeyoungeee/blog-feed-notifier/service"
	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	Service *service.SubscriptionService
}

func NewSubscriptionHandler(s *service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{Service: s}
}

func (h *SubscriptionHandler) GetUserSubscriptions(c *gin.Context) {
	userIDStr := c.Param("user_id")
	var userID uint
	_, err := fmt.Sscanf(userIDStr, "%d", &userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 유저 ID입니다"})
		return
	}

	subscriptions, err := h.Service.GetUserSubscriptions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]dto.GetSubscriptionResp, len(subscriptions))
	for i, sub := range subscriptions {
		resp[i] = dto.GetSubscriptionResp{
			ID:       sub.ID,
			BlogID:   sub.BlogID,
			BlogName: sub.Blog.Name,
		}
	}
	c.JSON(http.StatusOK, resp)
}
