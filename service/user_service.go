package service

import (
	"context"
	"errors"

	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"example.com/rest-api/repository"
	"example.com/rest-api/utils"
	"go.uber.org/zap"
)

type UserService struct {
	repo repository.Repository
}

func (ps *UserService) ValidateCredentials(ctx *context.Context, user *models.User) (string, error) {
	log := logger.Get(*ctx).With(zap.String("username", user.EmailId),
		zap.String("role", user.RoleId.String()),
		zap.String("method", "ValidateCredentials"))

	retrievedPassword, err := ps.repo.GetPassword(ctx, user)

	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	log.Info("reading user from db successful")
	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return "", errors.New("credentials invalid")
	}

	return ps.repo.GetRole(ctx, user.RoleId)
}

func (us *UserService) Save(ctx *context.Context, u *models.User) error {
	log := logger.Get(*ctx).With(
		zap.String("username", u.EmailId),
		zap.String("role", u.RoleId.String()),
		zap.String("method", "Save"))

	hashedPassword, err := utils.HashPassword(u.Password)
	u.Password = hashedPassword

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return us.repo.SaveUser(ctx, u)
}

// New creates new instance of Service
func NewUserService(repo repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}
