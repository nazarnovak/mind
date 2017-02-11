package data

import (
	"database/sql"
)

type Case struct {
	ID int
	CreatorId int
	DoctorId int
}

func (c *Case) Put() (int, error) {
	var id int

	row, err := DB.Query("INSERT INTO cases VALUES(DEFAULT, $1, DEFAULT)" +
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

func GetCaseById(cId int) (*Case, error) {
	c := Case{}
	err := DB.QueryRow("SELECT id, creatorid, doctorid FROM cases WHERE " +
		"id=$1;",
		cId).Scan(&c.ID, &c.CreatorId, &c.DoctorId)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &c, nil
	}
}

func GetCaseByIdCreatorId(caseId int, cId int) (*Case, error) {
	c := Case{}
	err := DB.QueryRow("SELECT id, creatorid, doctorid FROM cases WHERE " +
		"id=$1 AND creatorid=$2;",
		caseId, cId).Scan(&c.ID, &c.CreatorId, &c.DoctorId)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &c, nil
	}
}

func GetCaseByIdDoctorId(cId int, dId int) (*Case, error) {
	c := Case{}
	err := DB.QueryRow("SELECT id, creatorid, doctorid FROM cases WHERE " +
		"id=$1 AND doctorid=$2;",
		cId, dId).Scan(&c.ID, &c.CreatorId, &c.DoctorId)

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
	rows, err := DB.Query("SELECT id FROM cases WHERE creatorid=$1", cId)

	if err == sql.ErrNoRows {
		return []int{}, nil
	} else if err != nil {
		return []int{}, err
	}

	var id int
	cases := []int{}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return []int{}, err
		}
		cases = append(cases, id)
	}

	return cases, nil
}

func GetCasesByDoctorId(dId int) ([]int, error) {
	rows, err := DB.Query("SELECT id FROM cases WHERE doctorid=$1", dId)

	if err == sql.ErrNoRows {
		return []int{}, nil
	} else if err != nil {
		return []int{}, err
	}

	var id int
	cases := []int{}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return []int{}, err
		}
		cases = append(cases, id)
	}

	return cases, nil
}

func AssignDoctorToCase(doctorId int, caseId int) error {
	var rId int

	err := DB.QueryRow("UPDATE cases SET doctorid=$1 WHERE id=$2 " +
		"returning id;",
		doctorId, caseId).Scan(&rId)
	if err != nil || rId == 0 {
		return err
	}

	return nil
}

func GetPatientAlienCasesId(cId int) ([]int, error) {
	rows, err := DB.Query("SELECT id FROM cases WHERE creatorid!=$1", cId)

	if err == sql.ErrNoRows {
		return []int{}, nil
	} else if err != nil {
		return []int{}, err
	}

	var id int
	cases := []int{}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return []int{}, err
		}
		cases = append(cases, id)
	}

	return cases, nil
}

func GetDoctorAlienCasesId(cId int) ([]int, error) {
	rows, err := DB.Query("SELECT id FROM cases WHERE doctorid!=$1", cId)

	if err == sql.ErrNoRows {
		return []int{}, nil
	} else if err != nil {
		return []int{}, err
	}

	var id int
	cases := []int{}

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return []int{}, err
		}
		cases = append(cases, id)
	}

	return cases, nil
}
