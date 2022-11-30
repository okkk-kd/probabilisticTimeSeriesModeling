package users

type User struct {
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}

type CreateUser struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Authorization struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type AuthorizationResponse struct {
	SessionKey string `json:"session_key"`
}

type UpdateUserPassword struct {
	UserName        string `json:"user_name"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
