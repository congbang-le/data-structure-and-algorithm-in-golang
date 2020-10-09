package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
	prev  *node
}

type list struct {
	head *node
	tail *node
	size int
}

func (l *list) Append(input int) {
	n := &node{
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

func (l *list) Prepend(input int) {
	n := &node{
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

func (l *list) Insert(idx int, input int) {
	if l.size == 0 {
		l.Append(input)
	}
	n := &node{
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

func (l *list) Remove(idx int) {
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

func (l *list) Reverse() {
	if l.size <= 1 {
		return
	}
	temp := &node{}
	current := l.head
	for ; current != nil; current = current.prev {
		temp = current.prev
		current.prev = current.next
		current.next = temp
	}
	l.head = temp.prev
	return
}

func (l *list) ShowList() {
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
	l := &list{}
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
