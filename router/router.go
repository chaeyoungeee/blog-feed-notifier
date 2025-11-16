package router

import (
	"github.com/gin-gonic/gin"

	"github.com/chaeyoungeee/blog-feed-notifier/handler"
	"github.com/chaeyoungeee/blog-feed-notifier/repository"
	"github.com/chaeyoungeee/blog-feed-notifier/service"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.POST("/users", userHandler.CreateUser)

	return router
}
