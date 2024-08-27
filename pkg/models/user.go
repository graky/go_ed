package models

import (
	"database/sql"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) Create(db *sql.DB, hashedPassword string) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	return db.QueryRow(query, u.Email, hashedPassword).Scan(&u.ID)
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}