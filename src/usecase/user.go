package usecase

import (
	"application/src/model"
	"go.uber.org/zap"
)

type UserUseCaseInterface interface {
	GetAllUsers() ([]model.User, error)
	GetSingleUser(id int) (model.User, error)
	DeleteUser(id int) error
	CreateUser(user model.User) error
}

type UserRepoInterface interface {
	GetAllUsers() ([]model.User, error)
	GetSingleUser(id int) (model.User, error)
	DeleteUser(id int) error
	CreateUser(user model.User) error
}

type UserUseCase struct {
	repo   UserRepoInterface
	logger *zap.Logger
}

func (u UserUseCase) CreateUser(user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) GetAllUsers() ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) GetSingleUser(id int) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserUseCase(repo UserRepoInterface, logger *zap.Logger) UserUseCaseInterface {
	return &UserUseCase{
		repo:   repo,
		logger: logger.With(zap.String("type", "UserUseCase")),
	}
}
