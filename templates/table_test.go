package templates_test

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/pjherring/mysql-gen/def"
	"github.com/pjherring/mysql-gen/templates"
	"github.com/stretchr/testify/assert"
)

const expectedTableOutput = `
package users

func (t *User) Store() error {
	if t.IsStored {
		return t.insert()
	}

	return t.update()
}

func (t *User) update() error {
	_, err := gen.GetDb().Exec(
		"UPDATE users SET create_date = ?, group_id = ?, name = ?, telephone = ?, update_date = ? WHERE user_id = ?",
		t.CreateDate, t.GroupId, t.Name, t.Telephone, t.UpdateDate, t.UserId,
	)

	if err != nil {
		t.IsStored = true
	}

	return err
}

func (t *User) insert() error {
	r, err := gen.GetDb().Exec(
		"INSERT INTO users (create_date, group_id, name, telephone, update_date) VALUES (?, ?, ?, ?, ?)",
		t.CreateDate, t.GroupId, t.Name, t.Telephone, t.UpdateDate,
	)

	if err == nil {
		t.UserId = r.LastInsertId()
		t.IsStored = true
	}

	return err
}`

var expectedTableOutputLines []string = strings.Split(strings.TrimSpace(expectedTableOutput), "\n")

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
		"findById": "SELECT * FROM users WHERE user_id = ?",
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

	actualLines := strings.Split(strings.TrimSpace(b.String()), "\n")

	for i, p := range expectedTableOutputLines {
		if !assert.Equal(t, strings.TrimSpace(p), strings.TrimSpace(actualLines[i])) {
			return
		}
	}

}
