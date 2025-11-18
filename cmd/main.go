package main

import (
	"log"

	"github.com/chaeyoungeee/blog-feed-notifier/config"
	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"github.com/chaeyoungeee/blog-feed-notifier/handler"
	"github.com/chaeyoungeee/blog-feed-notifier/repository"
	"github.com/chaeyoungeee/blog-feed-notifier/router"
	"github.com/chaeyoungeee/blog-feed-notifier/scheduler"
	"github.com/chaeyoungeee/blog-feed-notifier/service"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected:", db)

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Blog{})
	db.AutoMigrate(&domain.Subscription{})

	userRepo := repository.NewUserRepo(db)
	blogRepo := repository.NewBlogRepo(db)
	subRepo := repository.NewSubscriptionRepo(db)

	userService := service.NewUserService(userRepo)
	blogService := service.NewBlogService(blogRepo)
	subService := service.NewSubscriptionService(subRepo)
	feedService := service.NewFeedService()

	userHandler := handler.NewUserHandler(userService)
	blogHandler := handler.NewBlogHandler(blogService)
	subscriptionHandler := handler.NewSubscriptionHandler(subService)

	scheduler := scheduler.NewScheduler(subService, feedService, blogService)
	scheduler.Start()

	r := router.NewRouter(userHandler, blogHandler, subscriptionHandler)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
