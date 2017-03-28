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