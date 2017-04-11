package main

import (
	"testing"
	"github.com/microlib/usermanager/lib"
	"errors"
	"context"
	"fmt"
	"reflect"
)

type FkUserManagerService struct{}

func (u *FkUserManagerService) FindUser(id string) (usermanager.UserInterface, error) {
	if id == "1" {
		return usermanager.NewUser("1", "Karl", "karl@email.com", "mypassword"), nil
	}
	return &usermanager.User{}, errors.New("User not found")
}

func TestMakeFindUserEndpoint(t *testing.T) {
	testCases := []struct {
		findUserReq findUserRequest
		expected findUserResponse
	}{
		{
			findUserRequest{Id:"1"},
			findUserResponse{
				Id:"1",
				Name: "Karl",
				Email: "karl@email.com",
				Err: nil,
			},
		},
		{
			findUserRequest{Id:"2"},
			findUserResponse{
				Id:"",
				Name: "",
				Email: "",
				Err: errors.New("User not found"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("User with id '%s'", tc.findUserReq.Id), func(t *testing.T) {
			res, _ := makeFindUserEndpoint(&FkUserManagerService{})(context.Background(), tc.findUserReq)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("got %s; want %s", res, tc.expected)
			}
		})
	}
}
