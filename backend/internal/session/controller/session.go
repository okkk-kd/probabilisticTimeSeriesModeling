package controller

import "probabilisticTimeSeriesModeling/internal/session/usecase"

type sessionCtrl struct {
	uc usecase.SessionUC
}

type SessionCtrl interface {
}

func NewSessionCtrl(uc usecase.SessionUC) (obj SessionCtrl, err error) {
	return &sessionCtrl{
		uc: uc,
	}, err
}
