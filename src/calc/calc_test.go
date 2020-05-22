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
