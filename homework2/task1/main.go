package main

import (
	"fmt"
)

func main() {

	var a, b uint
	fmt.Println("Расчет площади прямоугольника.")

	for true {
		fmt.Println("Введите размер сторон прямоугольника через пробел. Для завершения введите 0.")
		fmt.Scan(&a)
		if a == 0 {
			break
		} else {
			fmt.Scan(&b) // разделено с вводом "a", чтобы для выхода из цикла можно было ввести один 0. Но при этом обе переменные можно ввести в 1 потоке.
			fmt.Println("Площадь прямоугольника равна: ", a*b)
		}
	}

}
