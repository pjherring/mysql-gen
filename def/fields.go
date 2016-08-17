package def

type Fields []*Field

type FieldFilter func(f *Field) bool
type StringMember func(f *Field) string

func IsPrimaryKey(f *Field) bool {
	return f.IsPrimaryKey
}

func NotPrimaryKey(f *Field) bool {
	return !f.IsPrimaryKey
}

func (f Fields) Filter(filter FieldFilter) Fields {
	var retval Fields
	for _, field := range f {
		if filter(field) {
			retval = append(retval, field)
		}
	}

	return retval
}

func (f Fields) Strings(s StringMember) []string {
	var retval []string

	for _, field := range f {
		retval = append(retval, s(field))
	}

	return retval
}
