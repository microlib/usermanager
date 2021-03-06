package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeFindUserEndpoint(svc UserManagerServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findUserRequest)
		u, err := svc.FindUser(req.Id)
		return findUserResponse{
			Id: u.Id(),
			Name: u.Name(),
			Email: u.Email(),
			Err: err,
		}, nil
	}
}