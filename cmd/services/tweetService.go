package services

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/Sachin24704/go-backend/ent"
	"github.com/Sachin24704/go-backend/ent/tweet"
	"github.com/Sachin24704/go-backend/ent/user"
	"github.com/Sachin24704/go-backend/gen/app"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TweetService struct {
	client *ent.Client
}

func NewTweetService(client *ent.Client) *TweetService {
	return &TweetService{client: client}
}

// Add a new tweet
func (tweet *TweetService) CreateTweet(ctx context.Context, req *connect.Request[app.CreateTweetRequest]) (*connect.Response[app.Tweet], error) {
	var (
		client   = tweet.client
		tweetMsg = req.Msg.GetTweet()
		userId   = req.Msg.GetAuthor()
		tweetId  = uuid.NewString()
	)
	// to check if user_id is present in users table
	u, err := client.User.
		Query().
		Where(user.ID(userId)).
		Only(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist in DB"))
	}
	t, err := client.Tweet.Create().
		SetID(tweetId).SetTweet(tweetMsg).
		SetAuthor(u).Save(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("cannot create new tweet"))
	}
	res := connect.NewResponse(&app.Tweet{
		Id:         t.ID,
		CreateTime: t.CreatedAt.Format("02-01-2006 15:04:05"),
		Author:     t.UserID,
		Tweet:      t.Tweet,
	})
	res.Header().Set(userId, "New Tweet Created!!!")
	return res, nil
}

// listing tweets of a particular author
func (tweetService *TweetService) ListTweets(ctx context.Context, req *connect.Request[app.ListTweetsRequest]) (*connect.Response[app.ListTweetsResponse], error) {
	var (
		client = tweetService.client
		userId = req.Msg.GetAuthor()
	)
	// to check if user_id is present in users table
	_, err := client.User.
		Query().
		Where(user.ID(userId)).
		Only(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist in DB"))
	}
	t, err := client.Tweet.
		Query().
		Where(tweet.UserID(userId)).
		All(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("cannot query tweets"))
	}
	tweets := make([]*app.Tweet, 0)
	for _, v := range t {
		tweets = append(tweets, &app.Tweet{
			Id:         v.ID,
			Author:     v.UserID,
			Tweet:      v.Tweet,
			CreateTime: v.CreatedAt.Format("02-01-2006 15:04:05"),
		})

	}
	res := connect.NewResponse(&app.ListTweetsResponse{Tweets: tweets})
	return res, nil
}

// delete a tweet
func (tweet *TweetService) DeleteTweet(ctx context.Context, req *connect.Request[app.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	var (
		client  = tweet.client
		tweetId = req.Msg.GetId()
	)
	err := client.Tweet.DeleteOneID(tweetId).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("tweet_id doesnot exist in tweets"))
		}
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while deleting tweet from tweets"))
	}
	res := &connect.Response[emptypb.Empty]{}
	res.Header().Set(tweetId, "tweet deleted")
	return res, nil
}
