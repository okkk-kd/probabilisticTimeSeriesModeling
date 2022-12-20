package usecase

import (
	"probabilisticTimeSeriesModeling/internal/users"
	"probabilisticTimeSeriesModeling/internal/users/repository"
)

type userUC struct {
	repo repository.UserRepo
}

type UserUC interface {
	CreateUser(params users.CreateUser) (err error)
	UpdateUserPassword(params users.UpdateUserPassword) (err error)
	Authorization(params users.Authorization) (result users.AuthorizationResponse, err error)
}

func NewUserUC(repo repository.UserRepo) (obj UserUC, err error) {
	return &userUC{
		repo: repo,
	}, err
}

func (uc *userUC) CreateUser(params users.CreateUser) (err error) {
	err = uc.repo.CreateUser(params)
	if err != nil {
		return
	}
	return
}

func (uc *userUC) UpdateUserPassword(params users.UpdateUserPassword) (err error) {
	err = uc.repo.UpdateUserPassword(params)
	if err != nil {
		return
	}
	return
}

func (uc *userUC) Authorization(params users.Authorization) (result users.AuthorizationResponse, err error) {
	result, err = uc.repo.Authorization(params)
	if err != nil {
		return
	}
	return
}
