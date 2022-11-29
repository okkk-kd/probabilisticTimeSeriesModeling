package controller

import "probabilisticTimeSeriesModeling/internal/users/usecase"

type userCtrl struct {
	uc usecase.UserUC
}

type UserCtrl interface {
}

func NewUserCtrl(uc usecase.UserUC) (obj UserCtrl, err error) {
	return &userCtrl{
		uc: uc,
	}, err
}
