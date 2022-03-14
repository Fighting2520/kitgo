package service

import (
	"github.com/Fighting2520/kitgo/common/encryt"
	"github.com/Fighting2520/kitgo/common/errorx"
	"github.com/Fighting2520/kitgo/common/jwttoken"
	"github.com/pkg/errors"
)

type (
	IUserService interface {
		Login(request *LoginRequest) (*LoginResponse, error)
		Register(request *RegisterRequest) (*RegisterResponse, error)
		UserInfo(request *UserInfoRequest) (*UserInfoResponse, error)
	}

	LoginRequest struct {
		Username string `json:"username,optional"`
		Password string `json:"password,optional"`
	}

	LoginResponse struct {
		Token         string `json:"token"`
		ExpireSeconds int    `json:"expireSeconds"`
	}

	RegisterRequest struct {
	}

	RegisterResponse struct {
	}

	UserInfoRequest struct {
	}

	UserInfoResponse struct {
	}
)

func (s *Service) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := s.userModel.FindByUsername(req.Username)
	if err != nil {
		return nil, errorx.NewCodeError(2001, errors.Wrap(err, "find by username failed").Error())
	}
	if user.Password != encryt.Md5(req.Password) {
		return nil, errorx.NewCodeError(2002, errors.Wrap(err, "password check failed").Error())
	}
	tokenKey := "123456"
	expired := 86400
	token, err := jwttoken.GenerateToken(tokenKey, req.Username, user.Password, expired)
	if err != nil {
		return nil, errorx.NewCodeError(2003, errors.Wrap(err, "jwt token generate failed").Error())
	}
	return &LoginResponse{
		Token:         token,
		ExpireSeconds: expired,
	}, nil
}

func (s *Service) Register(request *RegisterRequest) (*RegisterResponse, error) {
	return &RegisterResponse{}, nil
}

func (s *Service) UserInfo(request *UserInfoRequest) (*UserInfoResponse, error) {
	return &UserInfoResponse{}, nil
}
