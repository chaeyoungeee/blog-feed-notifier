package handler

import (
	"net/http"

	"github.com/chaeyoungeee/blog-feed-notifier/dto"
	"github.com/chaeyoungeee/blog-feed-notifier/service"
	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	Service *service.BlogService
}

func NewBlogHandler(s *service.BlogService) *BlogHandler {
	return &BlogHandler{Service: s}
}

func (h *BlogHandler) GetBlogs(c *gin.Context) {
	blogs, err := h.Service.GetBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]dto.GetBlogResp, len(blogs))
	for i, blog := range blogs {
		resp[i] = dto.GetBlogResp{
			ID:      blog.ID,
			Name:    blog.Name,
			MainURL: blog.MainURL,
		}
	}
	c.JSON(http.StatusOK, resp)
}
