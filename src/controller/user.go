package controller

import (
	"application/src/model"
	"application/src/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserService struct {
	uc     usecase.UserUseCaseInterface
	logger *zap.Logger
}

func NewUserService(uc usecase.UserUseCaseInterface, logger *zap.Logger) *UserService {
	return &UserService{
		uc:     uc,
		logger: logger,
	}
}
func (u *UserService) CreateUser(c *gin.Context) {

	token, _ := usecase.GenerateToken(model.AuthUser{Username: "test"})
	c.JSON(http.StatusOK, gin.H{"access_token": token})

}
func (u *UserService) GetSingleUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetAllUsers(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) DeleteUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
