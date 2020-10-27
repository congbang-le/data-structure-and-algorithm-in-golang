package main

import (
	"fmt"
	"log"
)

type Node struct {
	Value string
	Next  *Node
}

type Stack struct {
	Head *Node
	Size int
}

func (s *Stack) Push(Input string) {
	NewNode := &Node{
		Value: Input,
	}
	if s.Head == nil {
		s.Head = NewNode
		s.Size++
		return
	}
	CurrentNode := s.Head
	s.Head = NewNode
	s.Head.Next = CurrentNode
	s.Size++
	return
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.Head = s.Head.Next
	s.Size--
	return
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return s.Head.Value
}

func (s *Stack) IsEmpty() bool {
	if s.Size == 0 && s.Head == nil {
		// fmt.Println("The stack is empty")
		return true
	}
	return false
}

func (s *Stack) ShowStack() {
	if s.IsEmpty() {
		return
	}
	CurrentNode := s.Head
	for ; CurrentNode != nil; CurrentNode = CurrentNode.Next {
		fmt.Printf(" <- %v", CurrentNode.Value)
	}
	return
}

func CheckBalancedParentheses(Input string) bool {
	if Input == "" {
		log.Fatalln("Empty string")
	}
	RoundBracketsStack := &Stack{}
	SquareBracketsStack := &Stack{}
	CurlyBracesStack := &Stack{}
	for _, v := range Input {
		Pivot := string(v)
		switch Pivot {
		case "(":
			RoundBracketsStack.Push(Pivot)
		case "[":
			SquareBracketsStack.Push(Pivot)
		case "{":
			CurlyBracesStack.Push(Pivot)
		case ")":
			RoundBracketsStack.Pop()
		case "]":
			SquareBracketsStack.Pop()
		case "}":
			CurlyBracesStack.Pop()
		}
	}
	if CurlyBracesStack.IsEmpty() && SquareBracketsStack.IsEmpty() && RoundBracketsStack.IsEmpty() {
		fmt.Println("This string has balanced parentheses")
		return true
	}
	fmt.Println("This string has not balanced parentheses")
	return false
}

func main() {
	s := &Stack{}
	fmt.Println("====Test Push====")
	s.Push("Amazon")
	s.Push("Netflix")
	s.Push("Google")
	s.Push("Apple")
	s.Push("Facebook")
	fmt.Println("Expect: <- Facebook <- Apple <- Google <- Netflix <- Amazon")
	fmt.Print("Got:")
	s.ShowStack()
	fmt.Println("")
	fmt.Println("====Test Pop====")
	s.Pop()
	s.Pop()
	fmt.Println("Expect: <- Google <- Netflix <- Amazon")
	fmt.Print("Got:")
	s.ShowStack()
	fmt.Println("")
	fmt.Println("====Test Peek====")
	fmt.Println("Expect: Google")
	fmt.Printf("Got: %v", s.Peek())
	fmt.Println("")
	Input := "[]0909()()(*)(*0)()()()({)}][[]()0P({9)9}{}]{}"
	CheckBalancedParentheses(Input)
}
