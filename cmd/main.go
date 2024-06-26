package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Sachin24704/go-backend/cmd/services"
	"github.com/Sachin24704/go-backend/ent"
	appconnect "github.com/Sachin24704/go-backend/gen/app/appconnect"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "postgres://postgres:123456@localhost:5432/twitter_orm?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//tweet := services.NewTweetService(client)
	//user := services.NewUserService(client)
	mux := http.NewServeMux()
	//user_path, user_handler := appconnect.NewUserServiceHandler(services.NewUserService(client))
	//tweet_path, tweet_handler := appconnect.NewTweetServiceHandler(services.NewTweetService(client))
	fmt.Println("DB and server started")
	mux.Handle(appconnect.NewUserServiceHandler(services.NewUserService(client.User)))
	mux.Handle(appconnect.NewTweetServiceHandler(services.NewTweetService(client)))
	log.Fatal(http.ListenAndServe(":8000", mux))
}
