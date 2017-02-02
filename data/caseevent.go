package data

import (
	"time"
	"os"
)

const (
	System   = 0
	Messages = 0
)

type CaseEvent struct {
	Id      int
	CaseId  int
	UserId  int
	TypeId  int
	Created string
	Content string
}

type Message struct {
	Id       string
	UserName string
	UserRole string
	Created  time.Time
	Content  string
}

func GreetMessage(id int) error {
	ce := CaseEvent{}
	ce.CaseId = id
	ce.UserId = System
	ce.TypeId = Messages
	ce.Created = time.Now().Format(time.RFC3339)
	ce.Content = os.Getenv("MIND_GREET")

	_, err := ce.Put()
	if err != nil {
		return err
	}

	return nil
}

func GetCaseEventsByCaseId(caseId string, since string) ([]Message, error) {
	ce := []Message{}

	q := `SELECT ce.id, u.name, u.role, ce.created, ce.content
		FROM case_events ce
		JOIN users u on ce.userid = u.id
		WHERE caseid=$1 AND typeid=$2 AND created > $3
		ORDER BY created`

	rows, err := DB.Query(q, caseId, Messages, since)

	if err != nil {
		return []Message{}, err
	}

	for rows.Next() {
		var id, un, ur, content string
		var created time.Time

		err = rows.Scan(&id, &un, &ur, &created, &content)
		if err != nil {
			return []Message{}, err
		}
		ce = append(ce, Message{id, un, ur, created, content})
	}

	if err != nil {
		return []Message{}, err
	}

	return ce, nil
}

func (ce *CaseEvent) Put() (int, error) {
	var id int

	row, err := DB.Query("INSERT INTO case_events VALUES(DEFAULT, $1, " +
		"$2, $3, $4, $5) returning id;", ce.CaseId, ce.UserId, ce.TypeId,
		ce.Created, ce.Content)
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
