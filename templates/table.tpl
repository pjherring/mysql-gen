package {{.Raw}}

func (t *{{.Name}}) Store() error {
	if t.IsStored {
		return t.insert()
	}

	return t.update()
}

func (t *{{.Name}}) update() error {
    _, err := gen.GetDb().Exec(
        "UPDATE {{.Raw}} SET {{.NonPrimaryKeyParams}} WHERE {{.PrimaryKeyParams}}",
        {{.NonPrimaryKeyNames}}, {{.PrimaryKeyNames}},
    )

    if err != nil {
        t.IsStored = true
    }

    return err
}
