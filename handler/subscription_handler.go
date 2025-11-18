package handler

import (
	"net/http"

	"github.com/chaeyoungeee/blog-feed-notifier/domain"
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
	userID, ok := ParseUserID(c)
	if !ok {
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

func (h *SubscriptionHandler) CreateSubscription(c *gin.Context) {
	userID, ok := ParseUserID(c)
	if !ok {
		return
	}

	var req dto.CreateSubscriptionReq

	if !BindJSON(c, &req) {
		return
	}

	subscription := &domain.Subscription{
		UserID: userID,
		BlogID: req.BlogID,
	}

	err := h.Service.CreateSubscription(subscription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *SubscriptionHandler) CreateSubscriptions(c *gin.Context) {
	userID, ok := ParseUserID(c)
	if !ok {
		return
	}

	var req dto.CreateSubscriptionsReq
	if !BindJSON(c, &req) {
		return
	}

	subscriptions := make([]*domain.Subscription, len(req.BlogIDs))
	for i, blogID := range req.BlogIDs {
		subscriptions[i] = &domain.Subscription{
			UserID: userID,
			BlogID: blogID,
		}
	}

	err := h.Service.CreateSubscriptions(userID, subscriptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *SubscriptionHandler) DeleteSubscription(c *gin.Context) {
	userID, ok := ParseUserID(c)
	if !ok {
		return
	}

	subscritpionID, ok := ParseSubscriptionID(c)
	if !ok {
		return
	}

	err := h.Service.DeleteSubscription(userID, subscritpionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
