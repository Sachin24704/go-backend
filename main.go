package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	app "github.com/Sachin24704/go-backend/gen/app"
	appconnect "github.com/Sachin24704/go-backend/gen/app/appconnect"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TwitterServer struct {
	db *sql.DB
}

// to create user
func (twitter *TwitterServer) CreateUser(ctx context.Context, req *connect.Request[app.CreateUserRequest]) (*connect.Response[app.User], error) {
	db := twitter.db
	name := req.Msg.GetName()
	user_id := uuid.NewString()
	_, err := db.Exec("INSERT INTO users (user_id, name) VALUES ($1, $2)", user_id, name)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	row := db.QueryRow("SELECT to_char(created_at, 'DD-MM-YYYY HH24:MI:SS') AS created_at FROM users WHERE user_id = ($1)", user_id)
	var createdAt string
	err = row.Scan(&createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeInvalidArgument, sql.ErrNoRows)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&app.User{
		Name:       name,
		Id:         user_id,
		CreateTime: createdAt,
	})
	return res, nil
}

// to delete user
func (twitter *TwitterServer) DeletUser(ctx context.Context, req *connect.Request[app.DeletUserRequest]) (*connect.Response[emptypb.Empty], error) {
	db := twitter.db
	user_id := req.Msg.GetId()
	_, err := db.Exec("DELETE FROM users WHERE user_id= ($1)", user_id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	// to delete tweets of user from tweets table
	_, err = db.Exec("DELETE FROM tweets WHERE user_id= ($1)", user_id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return nil, nil
}

// Add a new tweet
func (twitter *TwitterServer) CreateTweet(ctx context.Context, req *connect.Request[app.CreateTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	db := twitter.db
	tweet := req.Msg.GetTweet()
	user_id := req.Msg.GetAuthor()
	tweet_id := uuid.NewString()
	row := db.QueryRow("SELECT name FROM users WHERE user_id = ($1)", user_id)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeInvalidArgument, sql.ErrNoRows)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	_, err = db.Exec("INSERT INTO tweets (user_id, tweet ,tweet_id) VALUES ($1, $2, $3)", user_id, tweet, tweet_id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	var res connect.Response[emptypb.Empty]
	res.Header().Set(name, "New Tweet Created!!!")
	return nil, nil
}

// listing tweets of a particular author
func (twitter *TwitterServer) ListTweets(ctx context.Context, req *connect.Request[app.ListTweetsRequest]) (*connect.Response[app.ListTweetsResponse], error) {
	db := twitter.db
	user_id := req.Msg.GetAuthor()
	row := db.QueryRow("SELECT name FROM users WHERE user_id = ($1)", user_id)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("user doesnot exist")
			return nil, connect.NewError(connect.CodeInvalidArgument, sql.ErrNoRows)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	rows, err := db.Query("SELECT tweet_id, tweet, to_char(created_at, 'DD-MM-YYYY HH24:MI:SS') AS created_at FROM tweets WHERE user_id = $1", user_id)
	if err != nil {
		log.Println("cannot get tweets from table")
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	// defer rows.Close()
	tweets := make([]*app.Tweet, 0)
	for rows.Next() {
		var tweet, tweet_id, created_at string
		err := rows.Scan(&tweet_id, &tweet, &created_at)
		if err != nil {
			log.Println("cannot scan rows")
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		tweets = append(tweets, &app.Tweet{
			Id:         tweet_id,
			Author:     user_id,
			Tweet:      tweet,
			CreateTime: created_at,
		})
	}
	res := connect.NewResponse(&app.ListTweetsResponse{Tweets: tweets})
	return res, nil
}

// delete a tweet
func (twitter *TwitterServer) DeleteTweet(ctx context.Context, req *connect.Request[app.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	db := twitter.db
	tweet_id := req.Msg.GetId()
	_, err := db.Exec("DELETE FROM tweets WHERE tweet_id = ($1)", tweet_id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	var res *connect.Response[emptypb.Empty]
	res.Header().Set(tweet_id, "tweet deleted")
	return nil, nil
}

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
	twitter := &TwitterServer{db: db}
	mux := http.NewServeMux()
	user_path, user_handler := appconnect.NewUserServiceHandler(twitter)
	tweet_path, tweet_handler := appconnect.NewTweetServiceHandler(twitter)
	fmt.Println("user path : ", user_path)
	fmt.Println("tweet path : ", tweet_path)
	mux.Handle(user_path, user_handler)
	mux.Handle(tweet_path, tweet_handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
	// http.ListenAndServe(
	//     "localhost:8080",
	//     // Use h2c so we can serve HTTP/2 without TLS.
	//     h2c.NewHandler(mux, &http2.Server{}),
	// )
}
