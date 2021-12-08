package main

import (
	"errors"
	"fmt"
	"io"
)

type node struct {
	value    string
	nextNode *node
}

type list struct {
	firstNode *node
	lastNode  *node
}

func (list *list) isEmpty() bool {
	return list.firstNode == nil
}

func (list *list) push(value string) {
	newNode := node{value: value}
	if list.isEmpty() {
		list.firstNode = &newNode
		list.lastNode = &newNode
		return
	}
	list.lastNode.nextNode = &newNode
	list.lastNode = &newNode
}

// Разворачивает список в обратную сторону
func (list *list) reverse() {
	nodePointer := list.firstNode
	list.lastNode = nil

	for nodePointer != nil {
		nodePointer.nextNode, nodePointer, list.lastNode = list.lastNode, nodePointer.nextNode, nodePointer
	}
	list.firstNode, list.lastNode = list.lastNode, list.firstNode
}

func main() {
	fmt.Println("Введите строки для заполнения списка. Для прекращения ввода используйте EOF")
	var str string
	list := list{}
	for {
		if _, err := fmt.Scanln(&str); errors.Is(err, io.EOF) {
			break
		}
		list.push(str)
	}

	//fmt.Println("First: ", list.firstNode) // Выводит адрес list.firstNode.nextNode
	//fmt.Println("Last: ", list.lastNode)   // Выводит адрес list.lastNode.nextNode
	//fmt.Printf("First: %p\n", list.firstNode)
	//fmt.Printf("Last: %p\n", list.lastNode)

	// развернуть список
	list.reverse()

	fmt.Println("Развернутый список:")
	nodePointer := list.firstNode
	for nodePointer != nil {
		fmt.Println(nodePointer.value)
		nodePointer = nodePointer.nextNode
	}

	//fmt.Println("First: ", list.firstNode) // Выводит адрес list.firstNode.nextNode
	//fmt.Println("Last: ", list.lastNode)   // Выводит адрес list.lastNode.nextNode
	//fmt.Printf("First: %p\n", list.firstNode)
	//fmt.Printf("Last: %p\n", list.lastNode)
}
