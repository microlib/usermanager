package main

import (
	"github.com/go-kit/kit/log"
	//"time"
	"github.com/microlib/usermanager/lib"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UserManagerServiceInterface
}

func (mw loggingMiddleware) FindUser(user string) (output usermanager.UserInterface, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "users",
			"input", user,
			//"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.FindUser(user)
	return
}


