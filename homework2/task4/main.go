package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	var primes []int
	var steps int
	fmt.Println("Простые числа до числа N.")

	for true {
		primes = nil
		steps = 0
		fmt.Println("Введите число N. Для завершения введите 0.")
		fmt.Scanln(&n)
		if n == 0 {
			break
		} else {
			if n > 1 {
				primes = append(primes, 2) // 2 сразу добавляем как простое
				nSqrt := int(math.Sqrt(float64(n)))
				fmt.Println(nSqrt)
				for i := 3; i <= n; i++ {
					// нужно проверить, что числа до N не делятся на предыдущие простые числа
					isPrime := true
					for _, digit := range primes {
						steps += 1
						if digit > nSqrt { // оптимизация этого алгоритма: можно проверять делители до корня из N
							break
						} else if i%digit == 0 {
							isPrime = false
							break
						}
					}
					if isPrime {
						primes = append(primes, i)
					}
				}
			}

			fmt.Println("Простые числа: ", primes)
			fmt.Println("Шагов для вычисления: ", steps)
		}
	}

}
