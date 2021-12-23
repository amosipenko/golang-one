package fibb

import "errors"

var ErrInvalidNumber = errors.New("Введено неправильное число. Введите число >= 0")
var numbers = map[int]int{0: 0, 1: 1}

func CalcFibb(n int) (int, error) {
	if n < 0 {
		return 0, ErrInvalidNumber
	}
	return Fibb(n), nil
}

func Fibb(n int) int {
	_, exists := numbers[n]
	if !exists {
		numbers[n] = Fibb(n-1) + Fibb(n-2)
	}
	return numbers[n]
}
