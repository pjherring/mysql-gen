package def

import "strings"

type Query struct {
	Name         string
	IsMulti      bool
	Params       Fields
	SelectFields Fields
	Sql          string
}

//ParseQueriesFromMap give a map of name => sql, make query structs
func ParseQueries(raw rawDef) map[string]Query {

	retval := map[string]Query{}

	for name, sql := range raw.Queries {
		retval[name] = parseQueryFromSql(raw, name, sql)
	}

	return retval
}

//parseQueryFromSql given a query break it apart into select fields, params, etc...
func parseQueryFromSql(fields Fields, name, sql string) Query {
	return Query{
		Name:         strings.Title(name),
		IsMulti:      isMulti(name),
		Params:       parseQueryParams(sql, retval.fieldMap),
		SelectFields: parseSelectFields(fields, sql, retval.fieldMap),
		Sql:          sql,
	}
}

//isMulti determines if a query return multiple rows or not
func isMulti(sql string) bool {
	return strings.Contains(sql, "Many")
}

//parseQueryParams determines what fields the user needs to supply for input
func parseQueryParams(sql string, fields map[string]*Field) []*Field {

	retval := []*Field{}

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
				retval = append(retval, &Field{
					Arg:  "limit",
					Name: "Limit",
					Raw:  "LIMIT",
				})
				break
			} else if strings.ToUpper(parts[j]) == "OFFSET" {
				retval = append(retval, &Field{
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

func parseSelectFields(sql string, fields map[string]*Field) []*Field {

	selectFields := sql[len(SELECT):strings.Index(sql, FROM)]
	if strings.Contains(selectFields, "*") {
		return []*Field{}
	}

	var retval []*Field

	selectFieldParts := strings.Split(selectFields, ",")

	for _, part := range selectFieldParts {
		part = strings.TrimSpace(part)
		if f, ok := fields[part]; ok {
			retval = append(retval, f)
		}
	}

	return retval
}
