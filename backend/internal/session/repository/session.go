package repository

import (
	"github.com/jmoiron/sqlx"
	"probabilisticTimeSeriesModeling/config"
)

type sessionRepo struct {
	cfg  *config.Config
	pgDB *sqlx.DB
}

type SessionRepo interface {
}

func NewSessionRepo(cfg *config.Config, pgDB *sqlx.DB) (obj SessionRepo, err error) {
	return &sessionRepo{
		cfg:  cfg,
		pgDB: pgDB,
	}, err
}
