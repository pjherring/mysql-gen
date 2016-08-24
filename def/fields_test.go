package def_test

import (
	"sort"
	"testing"

	"github.com/pjherring/mysql-gen/def"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	fields := def.Fields{
		&def.Field{
			Name:         "one",
			IsPrimaryKey: true,
		},
		&def.Field{
			Name:         "two",
			IsPrimaryKey: false,
		},
	}

	filtered := fields.Filter(def.IsPrimaryKey)
	assert.Equal(t, 1, len(filtered))
	assert.Equal(t, "one", filtered[0].Name)
}

func TestStrings(t *testing.T) {

	fields := def.Fields{
		&def.Field{
			Name:         "one",
			IsPrimaryKey: true,
		},
		&def.Field{
			Name:         "two",
			IsPrimaryKey: false,
		},
	}

	s := fields.Strings(func(f *def.Field) string {
		return f.Name + " = ?"
	})
	assert.Equal(t, 2, len(s))
	assert.Equal(t, "one = ?", s[0])
	assert.Equal(t, "two = ?", s[1])
}

func TestSort(t *testing.T) {
	fields := def.Fields{
		&def.Field{
			Name: "one",
		},
		&def.Field{
			Name: "ross",
		},
		&def.Field{
			Name: "two",
		},
		&def.Field{
			Name: "ccc",
		},
	}

	sort.Sort(fields)
	assert.Equal(t, "ccc", fields[0].Name)
	assert.Equal(t, "one", fields[1].Name)
	assert.Equal(t, "ross", fields[2].Name)
	assert.Equal(t, "two", fields[3].Name)
}
