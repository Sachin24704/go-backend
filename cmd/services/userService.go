package services

import (
	"context"
	"database/sql"
	"errors"

	"connectrpc.com/connect"
	"github.com/Sachin24704/go-backend/gen/app"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// to create user
func (user *UserService) CreateUser(ctx context.Context, req *connect.Request[app.CreateUserRequest]) (*connect.Response[app.User], error) {
	var (
		db = user.db
	    name = req.Msg.GetName()
	    userId = uuid.NewString()
    )
	_, err := db.Exec("INSERT INTO users (user_id, name) VALUES ($1, $2)", userId, name)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("cannot insert into database"))
	}
	row := db.QueryRow("SELECT to_char(created_at, 'DD-MM-YYYY HH24:MI:SS') AS created_at FROM users WHERE user_id = ($1)", userId)
	var createdAt string
	err = row.Scan(&createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("cannot find the record in database"))
		}
		return nil, connect.NewError(connect.CodeInternal,errors.New("error while scanning row from query") )
	}
	res := connect.NewResponse(&app.User{
		Name:       name,
		Id:         userId,
		CreateTime: createdAt,
	})
	return res, nil
}

// to delete user
func (user *UserService) DeleteUser(ctx context.Context, req *connect.Request[app.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error) {
	var (
		db = user.db
	    userId = req.Msg.GetId()
    )
	row := db.QueryRow("SELECT name FROM users WHERE user_id = ($1)", userId)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist"))
		}
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while scanning for name from users"))
	}
	_, err = db.Exec("DELETE FROM users WHERE user_id= ($1)", userId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while deleting user from users"))
	}
	//to delete tweets of user from tweets table
	_, err = db.Exec("DELETE FROM tweets WHERE user_id= ($1)", userId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("error while deleting tweets of user"))
	}
	res := &connect.Response[emptypb.Empty]{}
	return res, nil
}
