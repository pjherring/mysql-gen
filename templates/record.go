package templates

import (
	"html/template"
	"io"

	"github.com/pjherring/mysql-gen/def"
)

var recordTpl *template.Template

func init() {
	var err error
	recordTpl, err = template.ParseFiles("record.tpl")
	template.Must(recordTpl, err)
}

func WriteRecord(w io.Writer, t def.Table) error {
	return recordTpl.Execute(w, newTemplateDot(t))
}
