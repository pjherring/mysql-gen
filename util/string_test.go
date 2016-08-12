package util_test

import (
	"testing"

	"github.com/pjherring/mysql-gen/util"
	"github.com/stretchr/testify/assert"
)

func TestUnderscoreToCamelCase(t *testing.T) {
	assert.Equal(t, "camelCase", util.UnderscoreToCamelCase("camel_case"))
}
