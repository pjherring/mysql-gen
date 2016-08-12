package def

import (
	"encoding/json"
	"strings"

	"github.com/pjherring/mysql-gen/util"
)

type Table struct {
	Name       string
	PrimaryKey Field
	Fields     map[string]Field
	Queries    map[string]Query
}

type Field struct {
	Name string
	Type string
	Raw  string
}

type Query struct {
	Name    string
	IsMulti bool
	Params  []Field
	Raw     string
}

type rawDef struct {
	Name    string            `json:"name"`
	Fields  map[string]string `json:"fields"`
	Queries map[string]string `json:"queries"`
	PK      string            `json:"primary_key"`
}

func ParseTable(b []byte) (Table, error) {

	var retval Table
	var raw rawDef

	if err := json.Unmarshal(b, &raw); err != nil {
		return retval, err
	}

	retval.Name = strings.Title(raw.Name)
	retval.Fields = map[string]Field{}

	for name, t := range raw.Fields {
		f := Field{
			Name: strings.Title(util.UnderscoreToCamelCase(name)),
			Type: t,
			Raw:  name,
		}

		retval.Fields[name] = f

		if name == raw.PK {
			retval.PrimaryKey = f
		}
	}

	return retval, nil
}
