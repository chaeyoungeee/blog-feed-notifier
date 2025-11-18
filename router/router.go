package router

import (
	"time"

	"github.com/chaeyoungeee/blog-feed-notifier/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *handler.UserHandler, blogHandler *handler.BlogHandler, subscriptionHandler *handler.SubscriptionHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api/v1")
	{
		api.POST("/users", userHandler.CreateUser)
		api.POST("/auth/login", userHandler.Login)
		api.GET("/blogs", blogHandler.GetBlogs)
		api.GET("/users/:user_id/subscriptions", subscriptionHandler.GetUserSubscriptions)
		api.POST("/users/:user_id/subscriptions", subscriptionHandler.CreateSubscription)
		api.DELETE("/users/:user_id/subscriptions/:subscription_id", subscriptionHandler.DeleteSubscription)
		api.POST("/users/:user_id/discord-webhook", userHandler.SetDiscordWebhook)
		api.POST("/users/:user_id/subscriptions/batch", subscriptionHandler.CreateSubscriptions)
	}

	return r
}
