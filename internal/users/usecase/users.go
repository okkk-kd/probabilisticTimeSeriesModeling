package usecase

import "probabilisticTimeSeriesModeling/internal/users/repository"

type userUC struct {
	repo repository.UserRepo
}

type UserUC interface {
}

func NewUserUC(repo repository.UserRepo) (obj UserUC, err error) {
	return &userUC{
		repo: repo,
	}, err
}
