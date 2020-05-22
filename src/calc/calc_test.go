package calc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOk(t *testing.T) {
	// GIVEN
	a := 5
	b := 4
	// WHEN
	result := Sum(a, b)
	// THEN
	assert.Equal(t, 9, result, fmt.Sprintf("Wrong sum!!"))
}

func TestDivOk(t *testing.T) {
	// GIVEN
	a := 6
	b := 2
	// WHEN
	result, err := Div(a, b)
	// THEN
	assert.Equal(t, 3, result, fmt.Sprintf("Wrong div!!"))
	assert.Nil(t, err, fmt.Sprintf("Wrong div - error not nil!!"))
}

func TestDivZero(t *testing.T) {
	// GIVEN
	a := 6
	b := 0
	// WHEN
	result, err := Div(a, b)
	// THEN
	assert.Equal(t, -1, result, fmt.Sprintf("Wrong div!!"))
	assert.NotNil(t, err, fmt.Sprintf("Wrong div - error nil!!"))
	assert.Equal(t, "can't work with 0", err.Error(), fmt.Sprintf("Wrong error message!!"))
}
