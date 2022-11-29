package usecase

import "probabilisticTimeSeriesModeling/internal/session/repository"

type sessionUC struct {
	repo repository.SessionRepo
}

type SessionUC interface {
}

func NewSessionUC(repo repository.SessionRepo) (obj SessionUC, err error) {
	return &sessionUC{
		repo: repo,
	}, err
}
