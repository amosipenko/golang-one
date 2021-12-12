package main

import (
	"fmt"
	"io"
)

type myByte []byte

func (myB myByte) Read(p []byte) (n int, err error) {
	var i int
	for i = 0; i < len(myB); i++ {
		p[i] = myB[i]
	}
	// Возврат err = nil уходит в бесконечный цикл, нужен io.EOF
	return i, io.EOF
}

func (myB *myByte) Write(p []byte) (n int, err error) {
	*myB = append(*myB, p...)
	return len(*myB), nil
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
