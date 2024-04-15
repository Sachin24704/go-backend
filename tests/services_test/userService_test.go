package services_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"github.com/Sachin24704/go-backend/gen/app"
	"github.com/Sachin24704/go-backend/gen/app/appconnect"
)

var userClient = appconnect.NewUserServiceClient(
	http.DefaultClient,
	"http://localhost:8000",
)

type tc struct {
	tcName        string
	userName      string
	userID        string
	expectedError error
}

var testCases = []tc{
	{tcName: "ValidName", userName: "dev", expectedError: nil},
	{tcName: "EmptyName", userName: "", expectedError: connect.NewError(connect.CodeNotFound, errors.New("empty name not allowed"))},
}

func TestCreateUser(t *testing.T) {
	for i, tc := range testCases {
		t.Run(tc.tcName, func(t *testing.T) {
			res, err := userClient.CreateUser(context.Background(), connect.NewRequest(&app.CreateUserRequest{Name: tc.userName}))
			if tc.expectedError == nil {
				if err != nil {
					t.Error("userId Exists test failed", err)
				}
			} else {
				if err.Error() != tc.expectedError.Error() {
					t.Error("UseridDoesnotExists test failed", err)
				}
			}
			if tc.expectedError == nil && err == nil {
				testCases[i].userID = res.Msg.Id
			}
		})
	}
}
func TestDeleteUser(t *testing.T) {
	testcases := []tc{
		{tcName: "ExistingUser", userID: testCases[0].userID, expectedError: nil},
		{tcName: "NonExistingUser", userID: "nonexistentuser", expectedError: connect.NewError(connect.CodeNotFound, errors.New("user_id doesnot exist in DB"))},
	}
	for _, tc := range testcases {
		t.Run(tc.tcName, func(t *testing.T) {
			_, err := userClient.DeleteUser(context.Background(), connect.NewRequest(&app.DeleteUserRequest{Id: tc.userID}))
			if tc.expectedError == nil {
				if err != nil {
					t.Error("userId Exists test failed", err)
				}
			} else {
				if err.Error() != tc.expectedError.Error() {
					t.Error("NonExcisting User test failed", err)
				}
			}
		})
	}
}
