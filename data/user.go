package data

import (
	"database/sql"
)

type User struct {
	ID int
	Name string
	Role int
}

func GetUserById(id string) (*User, error) {
	u := User{}

	err := DB.QueryRow("SELECT id, name, role FROM users WHERE id=$1",id).
		Scan(&u.ID, &u.Name, &u.Role)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &u, nil
	}
}
