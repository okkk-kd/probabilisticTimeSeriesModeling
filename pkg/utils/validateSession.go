package utils

import "github.com/jmoiron/sqlx"

func ValidateSession(db *sqlx.DB, sessionKey string) (authed bool, err error) {
	err = db.Get(&authed, "select case when authed = true then true else false end as authed from users.sessions where session_key = $1;", sessionKey)
	if err != nil {
		return
	}
	return
}
