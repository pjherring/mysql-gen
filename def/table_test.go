package def_test

import (
	. "testing"

	"github.com/pjherring/mysql-gen/def"
	"github.com/stretchr/testify/assert"
)

func TestParseTableFromFile(t *T) {

	json := []byte(`
    {
        "name": "users",
        "fields": {
            "user_id": "int64",
            "name": "string",
            "create_date": "mysql.NullTime",
            "update_date": "mysql.NullTime",
            "telephone": "string",
            "group_id": "int64"
        },
		"primary_key": "user_id",
        "queries": {
            "findById": "SELECT * FROM users WHERE user_id = ?",
            "findManyByGroupId": "SELECT * FROM users WHERE group_id = ? LIMIT ? OFFSET ?",
			"getManyUserIdsByName": "SELECT user_id FROM users WHERE name = ?"
        }
    }
	`)

	table, err := def.ParseTable(json)

	assert.Nil(t, err)

	assert.Equal(t, "Users", table.Name)
	assert.Equal(t, 6, len(table.Fields))

	assert.Equal(t, "UserId", table.Fields["user_id"].Name)
	assert.Equal(t, "int64", table.Fields["user_id"].Type)
	assert.Equal(t, "user_id", table.Fields["user_id"].Raw)

	assert.Equal(t, "Name", table.Fields["name"].Name)
	assert.Equal(t, "string", table.Fields["name"].Type)
	assert.Equal(t, "name", table.Fields["name"].Raw)

	assert.Equal(t, "CreateDate", table.Fields["create_date"].Name)
	assert.Equal(t, "mysql.NullTime", table.Fields["create_date"].Type)
	assert.Equal(t, "create_date", table.Fields["create_date"].Raw)

	assert.Equal(t, "UpdateDate", table.Fields["update_date"].Name)
	assert.Equal(t, "mysql.NullTime", table.Fields["update_date"].Type)
	assert.Equal(t, "update_date", table.Fields["update_date"].Raw)

	assert.Equal(t, "Telephone", table.Fields["telephone"].Name)
	assert.Equal(t, "string", table.Fields["telephone"].Type)
	assert.Equal(t, "telephone", table.Fields["telephone"].Raw)

	assert.Equal(t, "GroupId", table.Fields["group_id"].Name)
	assert.Equal(t, "int64", table.Fields["group_id"].Type)
	assert.Equal(t, "group_id", table.Fields["group_id"].Raw)

	assert.Equal(t, "UserId", table.PrimaryKey.Name)
	assert.Equal(t, "int64", table.PrimaryKey.Type)

	assert.Equal(t, "findById", table.Queries["findById"].Name)
	assert.False(t, table.Queries["findById"].IsMulti)
	assert.Equal(t, 1, len(table.Queries["findById"].Params))
	assert.Equal(t, 0, len(table.Queries["findById"].SelectFields))
	assert.Equal(t, "UserId", table.Queries["findById"].Params[0].Name)
	assert.Equal(t, "SELECT * FROM users WHERE user_id = ?", table.Queries["findById"].Sql)

	assert.Equal(t, "findManyByGroupId", table.Queries["findManyByGroupId"].Name)
	assert.True(t, table.Queries["findManyByGroupId"].IsMulti)
	assert.Equal(t, 3, len(table.Queries["findManyByGroupId"].Params))
	assert.Equal(t, 0, len(table.Queries["findById"].SelectFields))
	assert.Equal(t, "GroupId", table.Queries["findManyByGroupId"].Params[0].Name)
	assert.Equal(t, "Limit", table.Queries["findManyByGroupId"].Params[1].Name)
	assert.Equal(t, "limit", table.Queries["findManyByGroupId"].Params[1].Arg)
	assert.Equal(t, "Offset", table.Queries["findManyByGroupId"].Params[2].Name)
	assert.Equal(t, "offset", table.Queries["findManyByGroupId"].Params[2].Arg)

	//"getManyUserIdsByName: "SELECT user_id FROM users WHERE name = ?"
	assert.Equal(t, "getManyUserIdsByName", table.Queries["getManyUserIdsByName"].Name)
	assert.True(t, table.Queries["getManyUserIdsByName"].IsMulti)
	assert.Equal(t, 1, len(table.Queries["getManyUserIdsByName"].Params))
	assert.Equal(t, 1, len(table.Queries["getManyUserIdsByName"].SelectFields))
	assert.Equal(t, "UserId", table.Queries["getManyUserIdsByName"].SelectFields[0].Name)
	assert.Equal(t, "Name", table.Queries["getManyUserIdsByName"].Params[0].Name)

}
