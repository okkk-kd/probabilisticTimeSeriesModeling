package repository

import (
	"github.com/jmoiron/sqlx"
	"probabilisticTimeSeriesModeling/config"
)

type userRepo struct {
	cfg  *config.Config
	pgDB *sqlx.DB
}

type UserRepo interface {
}

func NewUserRepo(cfg *config.Config, pgDB *sqlx.DB) (obj UserRepo, err error) {
	return &userRepo{
		cfg:  cfg,
		pgDB: pgDB,
	}, err
}
