package services

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/Sachin24704/go-backend/gen/app"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TweetService struct {
	db *sql.DB
}

func NewTweetService(db *sql.DB) *TweetService{
	return &TweetService{db : db}
}

// Add a new tweet
func (tweet *TweetService) CreateTweet(ctx context.Context, req *connect.Request[app.CreateTweetRequest]) (*connect.Response[app.Tweet], error) {
	var(
		db = tweet.db
	    tweetMsg = req.Msg.GetTweet()
	    userId = req.Msg.GetAuthor()
	    tweetId = uuid.NewString()
    )
	// to check if user_id is present in users table
	row := db.QueryRow("SELECT user_id FROM users WHERE user_id = ($1)", userId)
	var userID string
	err := row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist in DB"))
		}
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while scanning for user_id"))
	}
	_, err = db.Exec("INSERT INTO tweets (user_id, tweet ,tweet_id) VALUES ($1, $2, $3)", userId, tweetMsg, tweetId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while creating tweet in DB"))
	}
	row = db.QueryRow("SELECT to_char(created_at, 'DD-MM-YYYY HH24:MI:SS') AS created_at FROM tweets WHERE tweet_id = ($1)", tweetId)
	var createdAt string
	err = row.Scan(&createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("tweet_id doesnot exist in tweets"))
		}
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while scanning for created_at"))
	}
	res := connect.NewResponse(&app.Tweet{
		Id: tweetId, 
		CreateTime: createdAt,
		Author: userId, 
		Tweet: tweetMsg,
	})
	res.Header().Set(userId, "New Tweet Created!!!")
	return res, nil
}

// listing tweets of a particular author
func (tweet *TweetService) ListTweets(ctx context.Context, req *connect.Request[app.ListTweetsRequest]) (*connect.Response[app.ListTweetsResponse], error) {
	var(
	db = tweet.db
	userId = req.Msg.GetAuthor()
)
	row := db.QueryRow("SELECT name FROM users WHERE user_id = ($1)", userId)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("user doesnot exist")
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist"))
		}
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while scanning for user_id in users table"))
	}
	rows, err := db.Query("SELECT tweet_id, tweet, to_char(created_at, 'DD-MM-YYYY HH24:MI:SS') AS created_at FROM tweets WHERE user_id = $1", userId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("cannot get tweets from table"))
	}
	tweets := make([]*app.Tweet, 0)
	for rows.Next() {
		var tweetMsg, tweetId, createdAt string
		err := rows.Scan(&tweetId, &tweetMsg, &createdAt)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, errors.New("cannot scan tweet from rows"))
		}
		tweets = append(tweets, &app.Tweet{
			Id:         tweetId,
			Author:     userId,
			Tweet:      tweetMsg,
			CreateTime: createdAt,
		})
	}
	res := connect.NewResponse(&app.ListTweetsResponse{Tweets: tweets})
	return res, nil
}

// delete a tweet
func (tweet *TweetService) DeleteTweet(ctx context.Context, req *connect.Request[app.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	var(
		db = tweet.db
		tweetId = req.Msg.GetId()
	)
	row := db.QueryRow("SELECT tweet_id FROM tweets WHERE tweet_id = ($1)", tweetId)
	var tweetID string
	err := row.Scan(&tweetID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("tweet_id doesnot exist in tweets"))
		}
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while scanning for tweet_id"))
	}
	
	_, err = db.Exec("DELETE FROM tweets WHERE tweet_id = ($1)", tweetId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("cannot delete the tweet from database"))
	}
	res := &connect.Response[emptypb.Empty]{}
	res.Header().Set(tweetId, "tweet deleted")
	return res, nil
}