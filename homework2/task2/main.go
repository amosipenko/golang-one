package main

import (
	"fmt"
	"math"
)

func main() {
	var s uint
	fmt.Println("Расчет диаметра и длины окружности.")

	for true {
		fmt.Println("Введите площадь круга. Для завершения введите 0.")
		fmt.Scanln(&s)
		if s == 0 {
			break
		}

		var d, l float64
		l = math.Sqrt(4 * float64(s) * math.Pi) // l = √4Sπ
		d = l / math.Pi                         // D = l/π
		fmt.Printf("Диаметр круга: %.2f. Длина окружности: %.2f\n", d, l)
	}
}
