package {{.Raw}}

type {{.Name}} struct {
{{range .Fields}}{{- .Name}} {{.Type}}
{{end -}}
    IsStored bool
}

func ({{.FirstInitial}} *{{.Name}}) Scan(s gen.ScanFunc) error {
    {{$i := .FirstInitial -}}
    return s(
        {{range .Fields}}&{{$i}}.{{.Name}},
        {{end -}}
    )
}
