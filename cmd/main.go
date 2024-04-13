package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Sachin24704/go-backend/cmd/services"
	appconnect "github.com/Sachin24704/go-backend/gen/app/appconnect"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost:5432/twitter_new?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connected")
	tweet := services.NewTweetService(db)
	user := services.NewUserService(db)
	mux := http.NewServeMux()
	user_path, user_handler := appconnect.NewUserServiceHandler(user)
	tweet_path, tweet_handler := appconnect.NewTweetServiceHandler(tweet)
	mux.Handle(user_path, user_handler)
	mux.Handle(tweet_path, tweet_handler)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
