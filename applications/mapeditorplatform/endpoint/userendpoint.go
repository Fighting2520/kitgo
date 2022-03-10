package endpoint

import (
	"context"
	"errors"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/log"
	"golang.org/x/time/rate"
)

type (
	UserEndPoint struct {
		LoginEndPoint    endpoint.Endpoint
		RegisterEndPoint endpoint.Endpoint
		InfoEndPoint     endpoint.Endpoint
	}
)

func NewUserEndPoint(s service.IService, logger log.Logger, limit *rate.Limiter) *UserEndPoint {
	return &UserEndPoint{
		LoginEndPoint:    Chain(LoggingMiddleware(logger), ratelimit.NewErroringLimiter(limit))(MakeUserLoginEndPoint(s)),
		RegisterEndPoint: Chain(LoggingMiddleware(logger), ratelimit.NewErroringLimiter(limit))(MakeUserRegisterEndPoint(s)),
		InfoEndPoint:     Chain(LoggingMiddleware(logger), ratelimit.NewErroringLimiter(limit))(MakeUserInfoEndPoint(s)),
	}
}

func MakeUserLoginEndPoint(s service.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		loginRequest, ok := request.(*service.LoginRequest)
		if !ok {
			return nil, errors.New("param type mistake")
		}
		return s.Login(loginRequest)
	}
}

func MakeUserRegisterEndPoint(s service.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}

func MakeUserInfoEndPoint(s service.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}
