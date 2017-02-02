package data

import (
	"database/sql"
)

type Case struct {
	ID int `json:"id"`
}

func (c *Case) Put() (int, error) {
	var id int

	row, err := DB.Query("INSERT INTO cases VALUES(DEFAULT) returning id;")
	if err != nil {
		return id, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func GetCaseById(id string) (*Case, error) {
	c := Case{}
	err := DB.QueryRow("SELECT id FROM cases WHERE id=$1", id).
		Scan(&c.ID)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &c, nil
	}
}
