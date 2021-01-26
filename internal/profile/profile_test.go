package profile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckToken(t *testing.T) {
	result := CheckToken("123456")
	assert.Nil(t, result)

	result = CheckToken("1234")
	assert.NotNil(t, result)

	result = CheckToken("1234")
	assert.Error(t, result, "The token 1234 must be composed by six digits")
}
