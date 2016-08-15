package templates

import (
	"html/template"
	"io"
	"strings"

	"github.com/pjherring/mysql-gen/def"
)

var recordTpl *template.Template

func init() {
	var err error
	recordTpl, err = template.ParseFiles("record.tpl")
	template.Must(recordTpl, err)
}

type recordTplDot struct {
	def.Table
	PackageName  string
	FirstInitial string
}

func WriteRecord(w io.Writer, t def.Table) error {
	return recordTpl.Execute(w, recordTplDot{
		Table:        t,
		FirstInitial: strings.ToLower(t.Name)[0:1],
	})
}
