package main

type findUserRequest struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type findUserResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Err error `json:"err,omitempty"`
}

func (f findUserResponse) error() error { return f.Err }

type findUsersRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type findUsersResponse struct {
	Users []*findUserResponse `json:"users"`
	Err error `json:"err,omitempty"`
}

func (f findUsersResponse) error() error { return f.Err }