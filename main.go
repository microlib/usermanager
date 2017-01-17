package main

import (
	"encoding/json"
	"log"
	"net/http"
	"context"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	ctx := context.Background()
	svc := UserManagerService{}

	findUserHandler := httptransport.NewServer(
		ctx,
		makeFindUserEndpoint(svc),
		decodeFindUserRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", findUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeFindUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request findUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

