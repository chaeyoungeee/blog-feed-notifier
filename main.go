package main

import (
	"log"

	"github.com/chaeyoungeee/blog-feed-notifier/config/db"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected:", db)
}
