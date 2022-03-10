package service

import (
	"github.com/go-kit/kit/log"
)

type Middleware func(next IService) IService

func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next IService) IService {
		for i := len(others) - 1; i >= 0; i-- { // reverse
			next = others[i](next)
		}
		return outer(next)
	}
}

type (
	loggingMiddleware struct {
		logger log.Logger
		next   IService
	}
)

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next IService) IService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) Login(request *LoginRequest) (*LoginResponse, error) {
	defer func() {
		l.logger.Log("method", "Login")
	}()
	return l.next.Login(request)
}

func (l loggingMiddleware) Register(request *RegisterRequest) (*RegisterResponse, error) {
	defer func() {
		l.logger.Log("method", "Register")
	}()
	return l.next.Register(request)
}

func (l loggingMiddleware) UserInfo(request *UserInfoRequest) (*UserInfoResponse, error) {
	defer func() {
		l.logger.Log("method", "UserInfo")
	}()
	return l.next.UserInfo(request)
}
