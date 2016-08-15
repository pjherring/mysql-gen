package {{.Raw}}

type {{.Name}} struct {
{{- range .Fields}}   
    {{.Name}} {{.Type}}
{{- end}}
}

func ({{.FirstInitial}} *{{.Name}}) Scan(s gen.ScanFunc) error {
    return s(
        {{$i := .FirstInitial}}{{range .Fields}}{{$i}}.{{.Name}},{{end}}
    )
}
