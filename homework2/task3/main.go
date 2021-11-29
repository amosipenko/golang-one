package main

import (
	"fmt"
)

func main() {
	var a uint
	fmt.Println("Расчет составных частей числа.")

	for true {
		fmt.Println("Введите трехзначное число. Для завершения введите 0.")
		fmt.Scanln(&a)
		if a == 0 {
			break
		} else if a < 100 || a > 999 {
			fmt.Println("Введено неправильное число.")
		} else {
			var h, d, u uint
			u = a % 10
			d = (a%100 - u) / 10
			h = (a - d*10 - u) / 100
			fmt.Printf("Сотен: %d, десятков: %d, единиц: %d\n", h, d, u)
		}
	}

}
