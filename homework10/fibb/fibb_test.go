package fibb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibb(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expeсted int
		err      error
	}{{
		name:     "normal case",
		input:    10,
		expeсted: 55,
		err:      nil,
	}, {
		name:     "normal case",
		input:    33,
		expeсted: 3524578,
		err:      nil,
	}, {
		name:     "invalid number",
		input:    -10,
		expeсted: 0,
		err:      ErrInvalidNumber,
	}}

	for _, cse := range testCases {
		fibbNumber, err := CalcFibb(cse.input)
		assert.Equal(t, cse.expeсted, fibbNumber, cse.name)
		assert.Equal(t, cse.err, err, cse.name)
	}
}
