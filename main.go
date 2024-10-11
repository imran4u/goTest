package main

import (
	"fmt"
	"strings"
)

//Linked list

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l Node) String() string {
	return fmt.Sprintf("%d ", l.value)
}

// Add at end
func (l *LinkedList) Add(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		for iterator := l.head; iterator != nil; iterator = iterator.next {
			if iterator.next == nil {
				iterator.next = node
				break
			}
		}
	}
}

// Add from beginning
func (l *LinkedList) AddFromStart(value int) {
	node := &Node{value: value}
	node.next = l.head
	l.head = node
}

func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}
	if l.head.value == value {
		l.head = l.head.next
		return
	}
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.next.value == value {
			iterator.next = iterator.next.next
			return
		}
	}
}

func (l LinkedList) PrintVertically() {
	node := l.head
	for node != nil {
		fmt.Println(node)
		node = node.next
	}
}

func (l LinkedList) Print() {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(iterator.String())
	}
	fmt.Println(sb.String())
}

func main() {

	// LinkedList := new(LinkedList) //LinkedList
	// LinkedList := LinkedList{head: nil}
	// LinkedList := LinkedList{}
	linkedList := LinkedList{}

	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Print()
}
