package model

type LoginModel struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Premium  bool   `db:"is_premium"`
	Verified bool   `db:"verified"`
}
