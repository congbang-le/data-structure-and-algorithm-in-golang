package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type SinglyLinkedList struct {
	head *Node
	size int
}

func (l *SinglyLinkedList) Insert(idx int, input int) {
	n := &Node{
		value: input,
	}
	if idx == 0 {
		l.Prepend(input)
		return
	}
	if idx == l.size {
		l.Append(input)
		return
	}
	current := l.head
	previous := l.head
	for i := 0; i < idx; i++ {
		previous = current
		current = current.next
	}
	temp := current
	current = n
	current.next = temp
	previous.next = current
	l.size++
	return
}

func (l *SinglyLinkedList) Remove(idx int) {
	if idx == 0 {
		l.head = l.head.next
		return
	}
	current := l.head
	previous := l.head
	for i := 0; i < idx; i++ {
		previous = current
		current = current.next
	}
	previous.next = current.next
	l.size--
	return

}

func (l *SinglyLinkedList) Append(input int) {
	n := &Node{
		value: input,
	}
	if l.head == nil {
		l.head = n
		l.size++
		return
	}
	current := l.head
	for {
		if current.next == nil {
			current.next = n
			l.size++
			return
		}
		current = current.next
	}
}

func (l *SinglyLinkedList) Prepend(input int) {
	n := &Node{
		value: input,
	}
	if l.head == nil {
		l.head = n
		l.size++
		return
	}
	current := l.head
	l.head = n
	l.head.next = current
	l.size++
	return
}

func (l *SinglyLinkedList) Reverse() {
	if l.size == 1 {
		return
	}
	current := l.head
	var previous *Node
	var next *Node
	for {
		if current != nil {
			next = current.next
			current.next = previous
			previous = current
			current = next
		} else {
			l.head = previous
			return
		}
	}
}

func (l *SinglyLinkedList) ShowList() {
	current := l.head
	for {
		if current == nil {
			break
		}
		fmt.Printf(" %v ->", current.value)
		current = current.next
	}
	return
}

func main() {
	l := &SinglyLinkedList{}
	l.Append(5)
	l.Append(6)
	fmt.Println("Expect: 5 -> 6 ->")
	fmt.Printf("Got:")
	l.ShowList()
	fmt.Println("\nExpect: 3 -> 4 -> 5 -> 6 ->")
	l.Prepend(4)
	l.Prepend(3)
	fmt.Printf("Got:")
	l.ShowList()
	l.Insert(2, 9)
	fmt.Println("\nExpect: 3 -> 4 -> 9 -> 5 -> 6 ->")
	fmt.Printf("Got:")
	l.ShowList()
	l.Remove(2)
	fmt.Println("\nExpect: 3 -> 4 -> 5 -> 6 ->")
	fmt.Printf("Got:")
	l.ShowList()
	l.Reverse()
	fmt.Println("\nExpect: 6 -> 5 -> 4 -> 3 ->")
	fmt.Printf("Got:")
	l.ShowList()
}
