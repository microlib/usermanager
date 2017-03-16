package main

import (
	"encoding/json"
	"net/http"
	"context"

	httptransport "github.com/go-kit/kit/transport/http"
	"os"
	gokitlog "github.com/go-kit/kit/log"
	"log"
)


func main() {
	var svc UserManagerServiceInterface
	svc = &UserManagerService{}

	//endpoint := makeFindUserEndpoint(svc)
	//endpoint = loggingMiddleware(gokitlog.NewContext(logger).With("method", "users"))(endpoint)

	logger := gokitlog.NewLogfmtLogger(os.Stderr)
	svc = loggingMiddleware{logger, svc}

	findUserHandler := httptransport.NewServer(
		makeFindUserEndpoint(svc),
		decodeFindUserRequest,
		encodeResponse,
	)

	http.Handle("/user", findUserHandler)
	log.Fatal(http.ListenAndServe(":8080", findUserHandler))
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

