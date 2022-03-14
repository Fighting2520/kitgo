package endpoint

import (
	"context"
	"github.com/Fighting2520/kitgo/common/log/logx"
	"github.com/go-kit/kit/endpoint"
	"time"
)

func Chain(outer endpoint.Middleware, others ...endpoint.Middleware) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		for i := len(others) - 1; i >= 0; i-- { // reverse
			next = others[i](next)
		}
		return outer(next)
	}
}

func LoggingMiddleware(logger logx.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Infof("transport_error: %s, took:%d", err, time.Since(begin))
			}(time.Now())
			return next(ctx, request)
		}
	}
}
