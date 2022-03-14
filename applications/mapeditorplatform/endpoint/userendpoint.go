package endpoint

import (
	"context"
	"errors"
	"github.com/Fighting2520/kitgo/applications/mapeditorplatform/service"
	"github.com/Fighting2520/kitgo/common/log/logx"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
)

type (
	UserEndPoint struct {
		LoginEndPoint    endpoint.Endpoint
		RegisterEndPoint endpoint.Endpoint
		InfoEndPoint     endpoint.Endpoint
	}
)

func NewUserEndPoint(s service.IService, limit ratelimit.Allower) *UserEndPoint {
	return &UserEndPoint{
		LoginEndPoint:    Chain(ratelimit.NewErroringLimiter(limit))(MakeUserLoginEndPoint(s)),
		RegisterEndPoint: Chain(ratelimit.NewErroringLimiter(limit))(MakeUserRegisterEndPoint(s)),
		InfoEndPoint:     Chain(ratelimit.NewErroringLimiter(limit))(MakeUserInfoEndPoint(s)),
	}
}

func MakeUserLoginEndPoint(s service.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		loginRequest, ok := request.(*service.LoginRequest)
		if !ok {
			logx.Errorf("param type mistake")
			return nil, errors.New("param type mistake")
		}
		logx.Errorf("request: %+v", loginRequest)
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
