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

func makeFindUsersEndpoint(svc UserManagerServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findUsersRequest)

		// create params based on the request
		params := map[string]string{}
		params["name"] = req.Name
		params["email"] = req.Email

		u, err := svc.FindUsers(params)
		return findUserResponse{
			Id: u.Id(),
			Name: u.Name(),
			Email: u.Email(),
			Err: err,
		}, nil
	}
}