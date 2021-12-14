package main

import (
	"fmt"
	"io"
	"math"
)

type myByte []byte

func (myB myByte) Read(p []byte) (n int, err error) {
	count := int(math.Min(float64(len(p)), float64(len(myB))))
	for i := 0; i < count; i++ {
		p[i] = myB[i]
	}
	// Возврат err = nil уходит в бесконечный цикл, нужен io.EOF
	return count, io.EOF
}

func (myB *myByte) Write(p []byte) (n int, err error) {
	*myB = append(*myB, p...)
	return len(p), nil
}

func main() {
	var myB myByte

	str := "abcdefgh"
	if _, err := io.WriteString(&myB, str); err != nil {
		fmt.Println("Не удалось записать строку в myByte: ", err)
		return
	}

	bytes, err := io.ReadAll(myB)
	if err != nil {
		fmt.Println("Не удалось прочитать строку из myByte: ", err)
		return
	}

	println("Прочитано: ", string(bytes))
}
