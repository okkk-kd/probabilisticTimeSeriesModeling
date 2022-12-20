package session

type Session struct {
	SessionKey string `json:"session_key" db:"session_key"`
	Authed     bool   `json:"authed" db:"authed"`
	UserID     int    `json:"user_id" db:"user_id"`
}
