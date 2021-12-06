package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите порядковый номер числа Фибоначчи")

	var n int
	fmt.Scanln(&n)
	if n < 0 {
		panic("Введите число >= 0")
	}

	numbers := map[int]int{0: 0, 1: 1}

	fmt.Println("Число Фибоначчи: ", fib(n, numbers))

}

func fib(n int, numbers map[int]int) int {
	_, exists := numbers[n]
	if !exists {
		numbers[n] = fib(n-1, numbers) + fib(n-2, numbers)
	}
	return numbers[n]
}
