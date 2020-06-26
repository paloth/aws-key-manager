package profile

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckToken(t *testing.T) {
	result := CheckToken("123456")
	assert.Nil(t, result)

	result = CheckToken("1234")
	assert.NotNil(t, result)

	result = CheckToken("1234")
	assert.Equal(t, result, fmt.Errorf("The token 1234 must be composed by six digits"), "error")
}

// func TestGetConfig(t *testing.T){
// 	ctrl := gomock.NewController
// 	home, err := "/Users/test", nil

// }
