package templates

import (
	"strings"

	"github.com/pjherring/mysql-gen/def"
)

type templateDot struct {
	def.Table
	PackageName  string
	FirstInitial string
}

func newTemplateDot(t def.Table) templateDot {
	return templateDot{
		Table:        t,
		FirstInitial: strings.ToLower(t.Name)[0:1],
	}
}

func (t templateDot) PrimaryKeyNames() string {
	return strings.Join(
		t.Fields.Filter(def.IsPrimaryKey).Strings(nameMember),
		", ",
	)
}

func (t templateDot) PrimaryKeyParams() string {
	return strings.Join(
		t.Fields.Filter(def.IsPrimaryKey).Strings(paramMember),
		", ",
	)
}

func (t templateDot) NonPrimaryKeyParams() string {
	return strings.Join(
		t.Fields.Filter(def.NotPrimaryKey).Strings(paramMember),
		", ",
	)
}

func (t templateDot) NonPrimaryKeyNames() string {
	return strings.Join(
		t.Fields.Filter(def.NotPrimaryKey).Strings(nameMember),
		", ",
	)
}

func nameMember(f *def.Field) string {
	return "t." + f.Name
}

func paramMember(f *def.Field) string {
	return f.Raw + " = ?"
}
