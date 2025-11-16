package handler

import (
	"fmt"
	"net/http"

	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"github.com/chaeyoungeee/blog-feed-notifier/dto"
	"github.com/chaeyoungeee/blog-feed-notifier/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserReq

	if !BindJson(c, &req) {
		return
	}

	user := &domain.User{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
	}

	if err := h.Service.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginReq

	if !BindJson(c, &req) {
		return
	}

	user, err := h.Service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resp := dto.LoginResp{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) SetDiscordWebhook(c *gin.Context) {
	userIDStr := c.Param("user_id")
	var userID uint
	_, err := fmt.Sscanf(userIDStr, "%d", &userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 유저 ID입니다"})
		return
	}

	var req dto.SetDiscordWebhookReq

	if !BindJson(c, &req) {
		return
	}

	err = h.Service.SetDiscordWebhook(userID, req.DiscordWebhookURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
