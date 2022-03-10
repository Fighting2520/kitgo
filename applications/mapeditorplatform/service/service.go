package service

import "github.com/Fighting2520/kitgo/applications/mapeditorplatform/model"

type (
	IService interface {
		IUserService
	}

	Service struct {
		userModel *model.UserModel
	}
)

func NewService(userModel *model.UserModel) *Service {
	return &Service{
		userModel: userModel,
	}
}
