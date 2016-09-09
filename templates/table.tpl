package {{.Raw}}

func (t *{{.Name}}) Store() (err error) {
	if t.IsStored {
		err = t.insert()
	} else {
        err = t.update()
    }

    if err == nil {
        t.IsStored = true
    }

    return
}

func (t *{{.Name}}) update() error {
    _, err := gen.GetDb().Exec(
        "UPDATE {{.Raw}} SET {{.NonPrimaryKeyColumnsPlaceholders}} WHERE {{.PrimaryKeyColumnPlaceholders}}",
        {{.NonPrimaryKeyNames}}, {{.PrimaryKeyNames}},
    )

    return err
}

func (t *{{.Name}}) insert() error {
    r, err := gen.GetDb().QueryRow(
        "INSERT INTO {{.Raw}} ({{.NotAutoGeneratedColumns}}) VALUES ({{.Placeholders .NotAutoGeneratedColumnCnt}})",
        {{.NotAutoGeneratedNames}},
    )

    if err == nil {
        t.{{.AutoGeneratedName}} = r.LastInsertId()
    }

    return err
}

//finders
func Find({{.PrimaryKeyTypedArguments}}) (*{{.Name}}, error) {
    r := gen.GetDb().Exec(
        "SELECT {{.AllColumns}} FROM users WHERE {{.PrimaryKeyColumnPlaceholders}}",
        {{.PrimaryKeyArguments}},
    )

    retval := new({{.Name}})
    if err := retval.Scan(r.Scan); err != nil {
        return nil, err
    }

    return retval, nil
}

