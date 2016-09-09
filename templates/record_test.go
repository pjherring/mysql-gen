package templates_test

import (
	"bytes"
	"testing"

	"github.com/pjherring/mysql-gen/def"
	"github.com/pjherring/mysql-gen/templates"
	"github.com/pjherring/mysql-gen/util"
	"github.com/stretchr/testify/assert"
)

const expectedForTestWriteRecord = `package users

type User struct {
	CreateDate mysql.NullTime
	Name string
	UserId int64
	IsStored bool
}

func (u *User) Scan(s gen.ScanFunc) error {
	return s(
		&u.CreateDate,
		&u.Name,
		&u.UserId,
	)
}`

const recordJson = `{
	"name": "users",
	"fields": {
		"user_id": "int64",
		"name": "string",
		"create_date": "mysql.NullTime"
	},
	"primary_keys": ["user_id"],
	"queries": {
		"findById": "SELECT * FROM users WHERE user_id = ?",
		"findManyByGroupId": "SELECT * FROM users WHERE group_id = ? LIMIT ? OFFSET ?",
		"getManyUserIdsByName": "SELECT user_id FROM users WHERE name = ?"
	}
}`

func TestWriteRecord(t *testing.T) {

	b := new(bytes.Buffer)

	tableDef, err := def.ParseTable([]byte(recordJson))
	assert.Nil(t, err)

	err = templates.WriteRecord(b, tableDef)

	assert.Nil(t, err)

	if !util.CompareLines(expectedForTestWriteRecord, b.String()) {
		t.FailNow()
	}
}
