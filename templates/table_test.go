package templates_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/pjherring/mysql-gen/def"
	"github.com/pjherring/mysql-gen/templates"
	"github.com/pjherring/mysql-gen/util"
	"github.com/stretchr/testify/assert"
)

const expectedTableOutput = `
package users

func (t *User) Store() (err error) {
	if t.IsStored {
		err = t.insert()
	} else {
		err = t.update()
	}
	
	if err == nil {
		t.IsStored = true
	}

	return
}

func (t *User) update() error {
	_, err := gen.GetDb().Exec(
		"UPDATE users SET create_date = ?, group_id = ?, name = ?, telephone = ?, update_date = ? WHERE user_id = ?",
		t.CreateDate, t.GroupId, t.Name, t.Telephone, t.UpdateDate, t.UserId,
	)

	return err
}

func (t *User) insert() error {
	r, err := gen.GetDb().Exec(
		"INSERT INTO users (create_date, group_id, name, telephone, update_date) VALUES (?, ?, ?, ?, ?)",
		t.CreateDate, t.GroupId, t.Name, t.Telephone, t.UpdateDate,
	)

	if err == nil {
		t.UserId = r.LastInsertId()
	}

	return err
}

//finders
func Find(userId int64) (*User, error) {
	r := gen.GetDb().QueryRow(
		"SELECT create_date, group_id, name, telephone, update_date, user_id FROM users WHERE user_id = ?",
		userId,
	)

	retval := new(User)
	if err := retval.Scan(r.Scan); err != nil {
		return nil, err
	}

	return retval, nil
}

func FindManyByGroupId(groupId int64, limit int, offset int) ([]*User, error) {
	r := gen.GetDb().Query(
		"SELECT create_date, group_id, name, telephone, update_date, user_id FROM users WHERE group_id = ? LIMIT ? OFFSET ?",
		groupId, limit, offset,
	)

	retval := make([]*User, limit)
	defer r.Close()
	for r.Next() {
		m := new(User)
		if err := m.Scan(r.Scan); err != nil {
			return nil, err
		}

		retval = append(retval, m)
	}

	return retval, nil
}
`

const tableJson = `{
	"name": "users",
	"fields": {
		"user_id": "int64",
		"name": "string",
		"create_date": "mysql.NullTime",
		"update_date": "mysql.NullTime",
		"telephone": "string",
		"group_id": "int64"
	},
	"primary_keys": ["user_id"],
	"auto_generated": ["user_id"],
	"queries": {
		"findManyByGroupId": "SELECT * FROM users WHERE group_id = ? LIMIT ? OFFSET ?",
		"getManyUserIdsByName": "SELECT user_id FROM users WHERE name = ?"
	}
}`

func TestWriteTable(t *testing.T) {

	tableDef, err := def.ParseTable([]byte(tableJson))
	assert.Nil(t, err)
	b := new(bytes.Buffer)

	err = templates.WriteTable(b, tableDef)
	if !assert.Nil(t, err) {
		log.Fatal(err.Error())
	}

	if !util.CompareLines(expectedTableOutput, b.String()) {
		t.FailNow()
	}

}
