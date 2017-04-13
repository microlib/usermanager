package main

import (
	"testing"
	"github.com/microlib/usermanager/lib"
)

func TestUsersToFindUserResponse(t *testing.T) {
	users := []usermanager.UserInterface{}
	users = append(users, usermanager.NewUser("1", "Piero", "piero@email.it", ""))

	res := usersToFindUserResponse(users)
	if res[1].Name != "Piero" {
		t.Errorf("The name of the user should be %s", "Piero")
	}
}
