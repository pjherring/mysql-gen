package templates

import (
	"bytes"
	"strings"
	"testing"

	"github.com/pjherring/mysql-gen/def"
	"github.com/pjherring/mysql-gen/templates"
	"github.com/stretchr/testify/assert"
)

const expectedForTestWriteRecord = `package users

type User struct {
	Name string
	SignUpDate mysql.NullTime
	UserId int64
	IsStored bool
}

func (u *User) Scan(s gen.ScanFunc) error {
	return s(
		&u.Name,
		&u.SignUpDate,
		&u.UserId,
	)
}`

func TestWriteRecord(t *testing.T) {

	b := new(bytes.Buffer)

	err := templates.WriteRecord(b, def.Table{
		Raw:  "users",
		Name: "User",
		Fields: map[string]def.Field{
			"user_id": def.Field{
				Name: "UserId",
				Raw:  "user_id",
				Arg:  "userId",
				Type: "int64",
			},
			"name": def.Field{
				Name: "Name",
				Raw:  "name",
				Arg:  "name",
				Type: "string",
			},
			"sign_up_date": def.Field{
				Name: "SignUpDate",
				Raw:  "sign_up_date",
				Arg:  "signUpDate",
				Type: "mysql.NullTime",
			},
		},
	})

	r := strings.NewReplacer("\t", "", "\n", "", " ", "")

	assert.Nil(t, err)
	assert.Equal(t, r.Replace(expectedForTestWriteRecord), r.Replace(b.String()))
}
