package main

import "github.com/microlib/usermanager/lib"

type UserManagerServiceInterface interface {
	FindUser(string) (usermanager.UserInterface, error)
	FindUsers(params map[string]string) ([]usermanager.UserInterface, error)
}

type UserManagerService struct{}

func (u *UserManagerService) FindUser(id string) (usermanager.UserInterface, error) {
	users := usermanager.NewUsers([]map[string]string{
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
			"id":       "3",
			"name":     "Roberto",
			"email":    "robbie@italianovero.it",
			"password": "passwordsupersicura",
		},
	})
	res, err := users.Get(id)
	return res, err
}

func (u *UserManagerService) FindUsers(params map[string]string) ([]usermanager.UserInterface, error) {
	users := usermanager.NewUsers([]map[string]string{
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
			"id":       "3",
			"name":     "Roberto",
			"email":    "robbie@italianovero.it",
			"password": "passwordsupersicura",
		},
	})
	return users.Find(params)
}