package planetsvc

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type ServiceMiddleware func(Service) Service

type instrmw struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	Service
}

func InstrumentationMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	countResult metrics.Histogram,
) ServiceMiddleware {
	return func(next Service) Service {
		return instrmw{requestCount, requestLatency, countResult, next}
	}
}

func (mw instrmw) GetPlanet(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getPlanet", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Service.GetPlanet(ctx, s)
	return
}
