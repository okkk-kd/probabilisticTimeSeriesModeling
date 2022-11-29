package users

type User struct {
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}
