package {{.PackageName}}

import(
)

type {{.Name}} struct {
    {{range .Fields}}
    {{.Name}} {{.Type}}
    {{end}}
}

func ({{.FirstInitial}} *{{.Name}}) Scan(s gen.ScanFunc) error {
    return s(
        {{range .Fields}}
        {{.FirstInitial}}.{{.Name}},
        {{end}}
    )
}
