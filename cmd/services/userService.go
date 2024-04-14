package services

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/Sachin24704/go-backend/ent"
	"github.com/Sachin24704/go-backend/gen/app"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	client *ent.UserClient
}

func NewUserService(client *ent.UserClient) *UserService {
	return &UserService{client: client}
}

// to create user
func (user *UserService) CreateUser(ctx context.Context, req *connect.Request[app.CreateUserRequest]) (*connect.Response[app.User], error) {
	var (
		client = user.client
		name   = req.Msg.GetName()
		userId = uuid.NewString()
	)
	User, err := client.
		Create().
		SetID(userId).
		SetName(name).
		Save(ctx)
	if err != nil {
		if ent.IsValidationError(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("empty name not allowed"))
		}
		return nil, connect.NewError(connect.CodeInternal, errors.New("cannot create new user"))
	}
	res := connect.NewResponse(&app.User{
		Name:       User.Name,
		Id:         User.ID,
		CreateTime: User.CreatedAt.Format("02-01-2006 15:04:05"),
	})
	return res, nil
}

// to delete user
func (user *UserService) DeleteUser(ctx context.Context, req *connect.Request[app.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error) {
	var (
		client = user.client
		userId = req.Msg.GetId()
	)
	err := client.
		DeleteOneID(userId).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist in DB"))
		}
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while deleting user from users"))
	}
	res := &connect.Response[emptypb.Empty]{}
	return res, nil
}
