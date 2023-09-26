package repository

import (
	"application/src/model"
	"application/src/usecase"
	"go.uber.org/zap"
)

type UserRepo struct {
	logger *zap.Logger
	ds     *DataSource
}

func (u UserRepo) CreateUser(user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) GetAllUsers() ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) GetSingleUser(id int) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(logger *zap.Logger, ds *DataSource) usecase.UserRepoInterface {
	return &UserRepo{
		logger: logger.With(zap.String("type", "UserRepo")),
		ds:     ds,
	}
}
