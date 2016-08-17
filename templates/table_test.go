package templates_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/pjherring/mysql-gen/def"
	"github.com/pjherring/mysql-gen/templates"
	"github.com/stretchr/testify/assert"
)

/*
func (u *User) insert() error {
	r, err := gen.GetDb().Exec(
		"INSERT INTO users (name, create_date, update_date, telephone, group_id) VALUES (?, ?, ?, ?, ?)",
		t.Name, t.CreateDate, t.UpdateDate, t.Telephone, t.GroupId
	)

	if err == nil {
		t.UserId = r.LastInsertId()
		t.IsStored = true
	}

	return err
}
*/

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
		"UPDATE users SET name = ?, create_date = ?, update_date = ?, telephone = ?, group_id = ? WHERE user_id = ?",
		t.Name, t.CreateDate, t.UpdateDate, t.Telephone, t.GroupId, t.UserId,
	)

	if err != nil {
		t.IsStored = true
	}

	return err
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
	assert.Nil(t, err)
	assert.Equal(
		t,
		templates.TemplateReplacer.Replace(expectedTableOutput),
		templates.TemplateReplacer.Replace(b.String()),
		fmt.Sprintf("%s (expected) != %s (actual)", expectedTableOutput, b.String()),
	)

}
