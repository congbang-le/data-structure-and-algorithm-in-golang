package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *DoublyLinkedList) Append(input int) {
	n := &Node{
		value: input,
	}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.size++
		return
	}
	temp := l.tail
	l.tail = n
	l.tail.prev = temp
	temp.next = l.tail
	l.size++
	return
}

func (l *DoublyLinkedList) Prepend(input int) {
	n := &Node{
		value: input,
	}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.size++
		return
	}
	temp := l.head
	l.head = n
	l.head.next = temp
	temp.prev = l.head
	l.size++
	return
}

func (l *DoublyLinkedList) Insert(idx int, input int) {
	if l.size == 0 {
		l.Append(input)
	}
	n := &Node{
		value: input,
	}
	current := l.head
	for i := 0; i < idx-1; i++ {
		current = current.next
	}
	temp := current.next
	current.next = n
	n.prev = current
	n.next = temp
	l.size++
	return
}

func (l *DoublyLinkedList) Remove(idx int) {
	if l.size == 1 {
		l.head = nil
		l.tail = nil
		return
	}
	current := l.head
	for i := 0; i < idx; i++ {
		current = current.next
	}
	temp := current
	current.prev.next = temp.next
	current.prev = temp.prev
	l.size--
	return
}

func (l *DoublyLinkedList) Reverse() {
	if l.size <= 1 {
		return
	}
	temp := &Node{}
	current := l.head
	for ; current != nil; current = current.prev {
		temp = current.prev
		current.prev = current.next
		current.next = temp
	}
	l.head = temp.prev
	return
}

func (l *DoublyLinkedList) ShowList() {
	current := l.head
	for {
		if current == nil {
			break
		}
		fmt.Printf(" %v <->", current.value)
		current = current.next
	}
	return
}

func main() {
	l := &DoublyLinkedList{}
	l.Append(5)
	l.Append(6)
	fmt.Println("Expect: 5 <-> 6 <->")
	fmt.Printf("Got:")
	l.ShowList()
	fmt.Println("\nExpect: 3 <-> 4 <-> 5 <-> 6 <->")
	l.Prepend(4)
	l.Prepend(3)
	fmt.Printf("Got:")
	l.ShowList()
	l.Insert(2, 9)
	fmt.Println("\nExpect: 3 <-> 4 <-> 9 <-> 5 <-> 6 <->")
	fmt.Printf("Got:")
	l.ShowList()
	l.Remove(2)
	fmt.Println("\nExpect: 3 <-> 4 <-> 5 <-> 6 <->")
	fmt.Printf("Got:")
	l.ShowList()
	l.Reverse()
	fmt.Println("\nExpect: 6 <-> 5 <-> 4 <-> 3 <->")
	fmt.Printf("Got:")
	l.ShowList()
}
