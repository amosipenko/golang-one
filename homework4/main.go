package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	slice := make([]int, 0, 5)

	fmt.Println("Введите массив чисел. Для прекращения ввода используйте EOF")
	for {
		var a int
		if _, err := fmt.Scanln(&a); errors.Is(err, io.EOF) {
			break
		}

		slice = append(slice, a)
	}

	fmt.Println("Before: ", slice)

	// Сортируем: начиная со 2 элемента сравниваем с предыдущим и сдвигаем его влево
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && slice[j-1] > slice[j]; j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}

	fmt.Println("After : ", slice)
}
