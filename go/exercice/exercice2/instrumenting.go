package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           OperationService
}

func (mw instrumentingMiddleware) Somme(nombre1 int64, nombre2 int64) (output int64, err error)  {
	defer func(begin time.Time) {
		lvs := []string{"method", "somme", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Somme(nombre1, nombre2)
	return
}