package main

import (
	"strings"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   OperationService
}

func (mw loggingMiddleware) Somme(nombre1 int64, nombre2 int64) (output int64, err error) {
	var inputs = strings.Join([]string{strconv.FormatInt(nombre1, 10), strconv.FormatInt(nombre2, 10)}, "+")
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "count",
			"input", inputs,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Somme(nombre1, nombre2)
	return
}