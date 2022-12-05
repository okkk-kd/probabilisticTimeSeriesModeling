package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"probabilisticTimeSeriesModeling/config"
	"probabilisticTimeSeriesModeling/internal/users"
)

type userRepo struct {
	cfg  *config.Config
	pgDB *sqlx.DB
}

type UserRepo interface {
	CreateUser(params users.CreateUser) (err error)
	UpdateUserPassword(params users.UpdateUserPassword) (err error)
	Authorization(params users.Authorization) (result users.AuthorizationResponse, err error)
}

func NewUserRepo(cfg *config.Config, pgDB *sqlx.DB) (obj UserRepo, err error) {
	return &userRepo{
		cfg:  cfg,
		pgDB: pgDB,
	}, err
}

func (repo *userRepo) CreateUser(params users.CreateUser) (err error) {
	_, err = repo.pgDB.Exec(queryCreateUser, params.UserName, params.Password)
	if err != nil {
		return
	}
	return
}

func (repo *userRepo) UpdateUserPassword(params users.UpdateUserPassword) (err error) {
	var ok bool
	err = repo.pgDB.Get(&ok, queryUpdatePassword, params.NewPassword, params.UserName, params.CurrentPassword)
	if err != nil {
		return errors.New("Current password isn't correct")
	}
	if !ok {
		return errors.New("Current password isn't correct")
	}
	return
}

func (repo *userRepo) Authorization(params users.Authorization) (result users.AuthorizationResponse, err error) {
	var authed bool
	tx, err := repo.pgDB.Begin()
	if err != nil {
		return
	}
	err = repo.pgDB.Get(&authed, queryCheckUserParams, params.UserName, params.Password)
	if err != nil {
		return
	}
	_, err = tx.Exec(queryUpdateSessionKey, params.UserName)
	if err != nil {
		errR := tx.Rollback()
		if errR != nil {
			return
		}
		return
	}
	sessionKey, err := uuid.NewUUID()
	if err != nil {
		return
	}
	_, err = tx.Exec(queryAuthorizationInsertSessionKey, sessionKey.String(), true, params.UserName)
	if err != nil {
		errR := tx.Rollback()
		if errR != nil {
			return
		}
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	result.SessionKey = sessionKey.String()
	return
}
