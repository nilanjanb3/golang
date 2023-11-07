package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) preadd(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *LinkedList) printAll() {
	head := l.head
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func main() {
	fmt.Println("Test")
	myList := LinkedList{}
	myList.preadd(10)
	myList.preadd(20)
	myList.preadd(30)

	myList.printAll()

}
