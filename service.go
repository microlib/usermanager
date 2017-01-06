package main

import "github.com/microlib/usermanager/lib"

type UserManagerServiceInterface interface {
	FindUser(string) usermanager.UserInterface
}

type UserManagerService struct{}

func (u *UserManagerService) FindUser(id string) usermanager.UserInterface {
	users := usermanager.Users{[]map[string]string{
		map[string]string{
			"id":       "1",
			"name":     "Padraig",
			"email":    "padraig@irish.ie",
			"password": "123abc",
		},
		map[string]string{
			"id":       "2",
			"name":     "Padraig",
			"email":    "pat@superirish.com",
			"password": "mypassword123!",
		},
		map[string]string{
			"id":       "2",
			"name":     "Roberto",
			"email":    "robbie@italianovero.it",
			"password": "passwordsupersicura",
		},
	}}
	res, _ := users.Get(id)
	return res
}