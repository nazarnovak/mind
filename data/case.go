package data

import (
	"database/sql"
	"strconv"
)

type Case struct {
	ID int
	CreatorId int
}

func (c *Case) Put() (int, error) {
	var id int

	row, err := DB.Query("INSERT INTO cases VALUES(DEFAULT, $1) " +
		"returning id;", c.CreatorId)
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

func GetCasesByCreatorId(cId int) ([]int, error) {
	cIdStr := strconv.Itoa(cId)

	row, err := DB.Query("SELECT id FROM cases WHERE creatorid=$1", cIdStr)

	if err == sql.ErrNoRows {
		return []int{}, nil
	} else if err != nil {
		return []int{}, err
	}

	var id int
	cases := []int{}

	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return []int{}, err
		}
		cases = append(cases, id)
	}

	return cases, nil
}