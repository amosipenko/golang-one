package partsOfNumber

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		name      string
		input     uint
		expeсtedH uint
		expectedD uint
		expectedU uint
		err       error
	}{
		{
			name:      "normal case",
			input:     693,
			expeсtedH: 6,
			expectedD: 9,
			expectedU: 3,
			err:       nil,
		}, {
			name:      "normal case",
			input:     100,
			expeсtedH: 1,
			expectedD: 0,
			expectedU: 0,
			err:       nil,
		}, {
			name:      "impossible number case",
			input:     50,
			expeсtedH: 0,
			expectedD: 0,
			expectedU: 0,
			err:       ErrImpossibleNumber,
		}, {
			name:      "impossible number case",
			input:     1005,
			expeсtedH: 0,
			expectedD: 0,
			expectedU: 0,
			err:       ErrImpossibleNumber,
		},
	}

	for _, cse := range testCases {
		t.Run(cse.name, func(t *testing.T) {
			h, d, u, err := Calc(cse.input)
			if !errors.Is(err, cse.err) {
				t.Errorf("Ожидали ошибку: %s, получили: %s", cse.err, err)
				return
			}

			if h != cse.expeсtedH || d != cse.expectedD || u != cse.expectedU {
				t.Errorf("Ожидали сотни/десятки/единицы: %d/%d/%d, получили: %d/%d/%d",
					cse.expeсtedH, cse.expectedD, cse.expectedU, h, d, u)
			}
		})
	}
}

func ExampleCalc() {
	var a, h, d, u uint
	a = 582
	h, d, u, err := Calc(a)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Printf("Сотен: %d, десятков: %d, единиц: %d\n", h, d, u)
	// Output:
	// Сотен: 5, десятков: 8, единиц: 2
}

// Вызов:
// go test
