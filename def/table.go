package def

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/pjherring/mysql-gen/util"
)

type Table struct {
	Name       string
	Raw        string
	PrimaryKey Field
	Fields     map[string]Field
	Queries    map[string]Query
}

type Field struct {
	Arg  string
	Name string
	Type string
	Raw  string
}

type Query struct {
	Name         string
	IsMulti      bool
	Params       []Field
	SelectFields []Field
	Sql          string
}

type rawDef struct {
	Name    string            `json:"name"`
	Fields  map[string]string `json:"fields"`
	Queries map[string]string `json:"queries"`
	PK      string            `json:"primary_key"`
}

func ParseTable(b []byte) (Table, error) {

	retval := Table{
		Fields:  map[string]Field{},
		Queries: map[string]Query{},
	}

	var raw rawDef

	if err := json.Unmarshal(b, &raw); err != nil {
		return retval, err
	}

	retval.Raw = raw.Name
	retval.Name = strings.Title(raw.Name)

	for name, t := range raw.Fields {
		f := Field{
			Arg:  util.UnderscoreToCamelCase(name),
			Type: t,
			Raw:  name,
		}

		f.Name = strings.Title(f.Arg)

		retval.Fields[name] = f

		if name == raw.PK {
			retval.PrimaryKey = f
		}
	}

	for name, sql := range raw.Queries {
		q := Query{
			Name:         name,
			IsMulti:      isMulti(name),
			Params:       parseQueryFields(sql, retval.Fields),
			SelectFields: parseSelectFields(sql, retval.Fields),
			Sql:          sql,
		}

		retval.Queries[name] = q
	}

	return retval, nil
}

func isMulti(sql string) bool {
	log.Println(sql)
	return strings.Contains(sql, "Many")
}

func parseQueryFields(sql string, fields map[string]Field) []Field {

	retval := []Field{}

	parts := strings.Split(sql, " ")
	for i, part := range parts {
		if part != "?" {
			continue
		}

		for j := i - 1; j > 0; j-- {
			if f, ok := fields[parts[j]]; ok {
				retval = append(retval, f)
				break
			} else if strings.ToUpper(parts[j]) == "LIMIT" {
				retval = append(retval, Field{
					Arg:  "limit",
					Name: "Limit",
					Raw:  "LIMIT",
				})
				break
			} else if strings.ToUpper(parts[j]) == "OFFSET" {
				retval = append(retval, Field{
					Arg:  "offset",
					Name: "Offset",
					Raw:  "OFFSET",
				})
				break
			}
		}
	}

	return retval
}

const SELECT = "SELECT "
const FROM = "FROM"

func parseSelectFields(sql string, fields map[string]Field) []Field {

	selectFields := sql[len(SELECT):strings.Index(sql, FROM)]
	if strings.Contains(selectFields, "*") {
		return []Field{}
	}

	var retval []Field

	selectFieldParts := strings.Split(selectFields, ",")

	for _, part := range selectFieldParts {
		part = strings.TrimSpace(part)
		if f, ok := fields[part]; ok {
			retval = append(retval, f)
		}
	}

	return retval
}
