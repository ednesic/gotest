package planetsvc

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) GetPlanet(ctx context.Context, s string) (p string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetPlanet", "id", s, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetPlanet(ctx, s)
}
