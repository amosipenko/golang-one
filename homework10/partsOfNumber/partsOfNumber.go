package partsOfNumber

import (
	"errors"
)

var ErrImpossibleNumber = errors.New("Введено неправильное число. Введите число от 100 до 999.")

func Calc(a uint) (uint, uint, uint, error) {
	if a < 100 || a > 999 {
		return 0, 0, 0, ErrImpossibleNumber
	} else {
		var h, d, u uint
		u = a % 10
		d = (a%100 - u) / 10
		h = (a - d*10 - u) / 100
		return h, d, u, nil
	}
}
