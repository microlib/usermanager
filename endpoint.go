package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/microlib/usermanager/lib"
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
		if req.Name != "" {
			params["name"] = req.Name
		}
		if req.Email != "" {
			params["email"] = req.Email
		}

		u, err := svc.FindUsers(params)

		return findUsersResponse{
			Users: usersToFindUserResponse(u),
			Err: err,
		}, nil
	}
}

func usersToFindUserResponse([]usermanager.UserInterface) []*findUserResponse {
	findUserReponses := []*findUserResponse{}
	findUserReponses = append(findUserReponses, &findUserResponse{})
	return findUserReponses
}

