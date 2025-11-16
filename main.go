package main

import (
	"log"

	"github.com/chaeyoungeee/blog-feed-notifier/config/db"
	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"github.com/chaeyoungeee/blog-feed-notifier/router"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected:", db)

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Blog{})
	db.AutoMigrate(&domain.Subscription{})

	r := router.NewRouter(db)
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
