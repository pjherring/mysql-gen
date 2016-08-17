package templates

import (
	"html/template"
	"io"

	"github.com/pjherring/mysql-gen/def"
)

var tableTpl *template.Template

func init() {
	var err error
	tableTpl, err = template.ParseFiles("table.tpl")
	template.Must(tableTpl, err)
}

func WriteTable(w io.Writer, t def.Table) error {
	return tableTpl.Execute(w, newTemplateDot(t))
}
