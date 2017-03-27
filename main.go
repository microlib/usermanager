package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	gokitlog "github.com/go-kit/kit/log"
)


func main() {
	var svc UserManagerServiceInterface
	svc = &UserManagerService{}

	logger := gokitlog.NewLogfmtLogger(os.Stderr)
	svc = loggingMiddleware{logger,svc}

	handler := MakeHTTPHandler(svc, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", ":8088")
		errs <- http.ListenAndServe(":8088", handler)
	}()

	logger.Log("exit", <-errs)
}
