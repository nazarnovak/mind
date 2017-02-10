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

	err := DB.QueryRow("SELECT id, name, role FROM users WHERE id=$1;",id).
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

func GetUserByName(name string) (*User, error) {
	u := User{}

	err := DB.QueryRow("SELECT id, name, role FROM users WHERE name=$1;",
		name).Scan(&u.ID, &u.Name, &u.Role)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &u, nil
	}
}

func GetRandomDoctorId() (int, error) {
	var id int

	err := DB.QueryRow("SELECT id FROM users WHERE role = 1 ORDER BY " +
		"random() LIMIT 1;").
		Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		return 0, nil
	case err != nil:
		return 0, err
	}

	return id, nil
}
