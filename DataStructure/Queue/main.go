package main

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

func (q *Queue) Enqueue(Input interface{}) {
	NewNode := &Node{
		Value: Input,
	}
	if q.Head == nil {
		q.Head = NewNode
		q.Tail = NewNode
		q.Size++
		return
	}
	Current := q.Head
	q.Head = NewNode
	q.Head.Next = Current
	q.Size++
	return
}

func (q *Queue) Dequeue() {
	if q.Head == nil {
		fmt.Println("The queue is empty")
		return
	}
	Current := q.Head
	for {
		if Current.Next.Next == nil {
			Current.Next = nil
			q.Size--
			return
		} else {
			Current = Current.Next
		}
	}
}

func (q *Queue) Peek() interface{} {
	if q.Head == nil {
		fmt.Println("The queue is empty")
		return nil
	}
	Current := q.Head
	for ; Current.Next != nil; Current = Current.Next {
	}
	return Current.Value

}

func (q *Queue) ShowQueue() {
	if q.Head == nil {
		fmt.Println("The queue is empty")
		return
	}
	Current := q.Head
	for ; Current != nil; Current = Current.Next {
		fmt.Printf("%v -> ", Current.Value)
	}
	return
}

func main() {
	q := &Queue{}
	q.Enqueue("John")
	q.Enqueue("Jack")
	q.Enqueue("Mark")
	q.Enqueue("Tom")
	fmt.Println("Expect: Tom -> Mark -> Jack -> John ->")
	fmt.Print("Got: ")
	q.ShowQueue()
	fmt.Println("")
	q.Dequeue()
	q.Dequeue()
	fmt.Println("Expect: Tom -> Mark -> ")
	fmt.Print("Got: ")
	q.ShowQueue()
	fmt.Println("")
	fmt.Println("Expect: Mark")
	fmt.Printf("Got: %v\n", q.Peek())

}
