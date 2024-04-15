package services_test

import (
	"context"
	"errors"
	"log"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"github.com/Sachin24704/go-backend/gen/app"
	"github.com/Sachin24704/go-backend/gen/app/appconnect"
)

var tweetClient = appconnect.NewTweetServiceClient(
	http.DefaultClient,
	"http://localhost:8000",
)

func TestCreateTweet(t *testing.T) {
	//create a user
	user, err := userClient.CreateUser(context.Background(), connect.NewRequest(&app.CreateUserRequest{Name: "test_user"}))
	if err != nil {
		t.Fatal("cannot create a test user for creating tweets", err)
	}
	userID := user.Msg.Id
	testcases := []struct {
		tcname        string
		userid        string
		tweet         string
		expectedError error
	}{
		{tcname: "UseridExists", userid: userID, tweet: "test", expectedError: nil},
		{tcname: "UseridDoesNotExists", userid: "nonexistentuser", tweet: "test", expectedError: connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist in DB"))},
	}
	for _, tc := range testcases {
		t.Run(tc.tcname, func(t *testing.T) {
			_, err := tweetClient.CreateTweet(context.Background(), connect.NewRequest(&app.CreateTweetRequest{Tweet: tc.tweet, Author: tc.userid}))
			if tc.expectedError == nil {
				if err != nil {
					t.Error("userId Exists test failed", err)
				}
			} else {
				if err.Error() != tc.expectedError.Error() {
					t.Error("UseridDoesnotExists test failed", err)
				}
			}
		})
	}
	t.Cleanup(func() {
		_, err := userClient.DeleteUser(context.Background(), connect.NewRequest(&app.DeleteUserRequest{Id: userID}))
		if err != nil {
			log.Println("error in cleaning up the resources")
		}
	})
}

func TestListTweets(t *testing.T) {
	//create a user
	user, err := userClient.CreateUser(context.Background(), connect.NewRequest(&app.CreateUserRequest{Name: "test_user"}))
	if err != nil {
		t.Fatal("cannot create a test user for creating tweets", err)
	}
	userID := user.Msg.Id
	t.Cleanup(func() {
		_, err := userClient.DeleteUser(context.Background(), connect.NewRequest(&app.DeleteUserRequest{Id: userID}))
		if err != nil {
			log.Println("error in cleaning up the resources")
		}
	})
	testcases := []struct {
		tcname        string
		userid        string
		tweet         string
		expectedError error
	}{
		{tcname: "UseridExists", userid: userID, tweet: "test", expectedError: nil},
		{tcname: "UseridDoesNotExists", userid: "nonexistentuser", tweet: "test", expectedError: connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist in DB"))},
	}

	for _, tc := range testcases {
		t.Run(tc.tcname, func(t *testing.T) {
			_, err := tweetClient.ListTweets(context.Background(), connect.NewRequest(&app.ListTweetsRequest{Author: tc.userid}))
			if tc.expectedError == nil {
				if err != nil {
					t.Error("userId Exists test failed", err)
				}
			} else {
				if err.Error() != tc.expectedError.Error() {
					t.Error("UseridDoesnotExists test failed", err)
				}
			}
		})
	}
}

func TestDeleteTweet(t *testing.T) {
	//create a user
	user, err := userClient.CreateUser(context.Background(), connect.NewRequest(&app.CreateUserRequest{Name: "test_user"}))
	if err != nil {
		t.Fatal("cannot create a test user for creating tweets", err)
	}
	userID := user.Msg.Id
	t.Cleanup(func() {
		_, err := userClient.DeleteUser(context.Background(), connect.NewRequest(&app.DeleteUserRequest{Id: userID}))
		if err != nil {
			log.Println("error in cleaning up the resources")
		}
	})
	tweet, err := tweetClient.CreateTweet(context.Background(), connect.NewRequest(&app.CreateTweetRequest{Tweet: "test", Author: userID}))
	if err != nil {
		t.Fatal("cannot create a test tweet", err)
	}
	tweetID := tweet.Msg.Id
	testcases := []struct {
		tcname        string
		tweetId       string
		expectedError error
	}{
		{tcname: "TweetidExists", tweetId: tweetID, expectedError: nil},
		{tcname: "TweetidDoesNotExists", tweetId: "nonexistentTweetId", expectedError: connect.NewError(connect.CodeNotFound, errors.New("tweet_id doesnot exist in tweets"))},
	}
	for _, tc := range testcases {
		t.Run(tc.tcname, func(t *testing.T) {
			_, err := tweetClient.DeleteTweet(context.Background(), connect.NewRequest(&app.DeleteTweetRequest{Id: tc.tweetId}))
			if tc.expectedError == nil {
				if err != nil {
					t.Error("TweetId Exists test failed", err)
				}
			} else {
				if err.Error() != tc.expectedError.Error() {
					t.Error("TweetidDoesNotExists test failed", err)
				}
			}
		})
	}
}
