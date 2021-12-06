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

	fmt.Println("First: ", list.firstNode)
	fmt.Println("Last: ", list.lastNode)

	// развернуть список
	fmt.Println("Развернутый список:")
	nodePointer := list.firstNode
	for nodePointer != nil {
		fmt.Println(nodePointer.value)
		nodePointer = nodePointer.nextNode
	}
}
