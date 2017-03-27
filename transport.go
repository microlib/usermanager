package main

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"encoding/json"
	"context"
	"bytes"
	"io/ioutil"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHTTPHandler(s UserManagerServiceInterface, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	//e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// GET     /user/:id			retrieves the given user by id
	// GET     /user                       	retrieve all users given the search parameters

	// try it with this: curl -X GET localhost:8088/user/2
	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		makeFindUserEndpoint(s),
		decodeFindUserRequest,
		encodeResponse,
		options...,
	))

	//r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
	//	e.GetProfileEndpoint,
	//	decodeGetProfileRequest,
	//	encodeResponse,
	//	options...,
	//))

	return r
}

func decodeFindUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request findUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return findUserRequest{Id: id}, nil
}


// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encodeRequest likewise JSON-encodes the request to the HTTP request body.
// Don't use it directly as a transport/http.Client EncodeRequestFunc:
// profilesvc endpoints require mutating the HTTP method and request path.
func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}


func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	//case ErrNotFound:
	//	return http.StatusNotFound
	//case ErrAlreadyExists, ErrInconsistentIDs:
	//	return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}