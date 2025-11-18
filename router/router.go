package router

import (
	"github.com/gin-gonic/gin"

	"github.com/chaeyoungeee/blog-feed-notifier/handler"
)

func NewRouter(userHandler *handler.UserHandler, blogHandler *handler.BlogHandler, subscriptionHandler *handler.SubscriptionHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/users", userHandler.CreateUser)
		api.POST("/auth/login", userHandler.Login)
		api.GET("/blogs", blogHandler.GetBlogs)
		api.GET("/users/:user_id/subscriptions", subscriptionHandler.GetUserSubscriptions)
		api.POST("/users/:user_id/subscriptions", subscriptionHandler.CreateSubscription)
		api.POST("/users/:user_id/discord-webhook", userHandler.SetDiscordWebhook)
		api.POST("/users/:user_id/subscriptions/batch", subscriptionHandler.CreateSubscriptions)
	}

	return router
}
