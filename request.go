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
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}
