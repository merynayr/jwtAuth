package model

// User представляет пользователя системы
type User struct {
	ID string `json:"id,omitempty" db:"user_id"`
}
