package util_test

import (
	"strings"
	"testing"

	"github.com/pjherring/mysql-gen/util"
	"github.com/stretchr/testify/assert"
)

func TestStringParts(t *testing.T) {
	s := util.StringParts{"one", "two", "three"}
	s.Each(strings.ToUpper)
	assert.Equal(t, "ONE", s[0])
	assert.Equal(t, "TWO", s[1])
	assert.Equal(t, "THREE", s[2])
}
