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

	findUsersHandler := httptransport.NewServer(
		makeFindUserEndpoint(svc),
		decodeFindUserRequest,
		encodeResponse,
	)

	http.Handle("/user", findUsersHandler)
	http.Handle("/user", findUsersHandler)
	log.Fatal(http.ListenAndServe(":8080", findUsersHandler))
}
