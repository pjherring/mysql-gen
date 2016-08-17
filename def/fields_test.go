package def_test

import (
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
