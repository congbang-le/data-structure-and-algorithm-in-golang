package main

import (
	"fmt"
)

type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (Tree *BinarySearchTree) Insert(Input int) {
	if Tree.Root == nil {
		Tree.Root = &Node{
			Value: Input,
		}
	} else {
		Tree.Root.Insert(Input)
	}
}

func (N *Node) Insert(Input int) {
	NewNode := &Node{
		Value: Input,
	}
	if Input < N.Value {
		if N.LeftChild == nil {
			N.LeftChild = NewNode
		} else {
			N.LeftChild.Insert(Input)
		}
	} else {
		if N.RightChild == nil {
			N.RightChild = NewNode
		} else {
			N.RightChild.Insert(Input)
		}
	}
}

func LookUp(Root *Node, Input int) bool {
	if Root == nil {
		return false
	}
	if Root.Value == Input {
		return true
	}
	if Input < Root.Value {
		if LookUp(Root.LeftChild, Input) {
			return true
		}
	} else {
		if LookUp(Root.RightChild, Input) {
			return true
		}
	}
	return false
}

func Remove() {

}

func ShowTreeInOrder(N *Node) {
	if N == nil {
		return
	} else {
		ShowTreeInOrder(N.LeftChild)
		fmt.Printf(" %v ", N.Value)
		ShowTreeInOrder(N.RightChild)
	}
}

func ShowTreePreOrder(N *Node) {
	if N == nil {
		return
	} else {
		fmt.Printf(" %v ", N.Value)
		ShowTreePreOrder(N.LeftChild)
		ShowTreePreOrder(N.RightChild)
	}
}

func ShowTreePostOrder(N *Node) {
	if N == nil {
		return
	} else {
		ShowTreePostOrder(N.LeftChild)
		ShowTreePostOrder(N.RightChild)
		fmt.Printf(" %v ", N.Value)
	}
}

func main() {
	var Tree BinarySearchTree
	Tree.Insert(10)
	Tree.Insert(8)
	Tree.Insert(9)
	Tree.Insert(11)
	Tree.Insert(6)
	Tree.Insert(14)
	Tree.Insert(15)
	Tree.Insert(12)
	fmt.Println("In Order: ")
	ShowTreeInOrder(Tree.Root)
	fmt.Println("\nPost Order: ")
	ShowTreePostOrder(Tree.Root)
	fmt.Println("\nPre Order: ")
	ShowTreePreOrder(Tree.Root)
	fmt.Println()
	fmt.Println(LookUp(Tree.Root, 5))
	fmt.Println(LookUp(Tree.Root, 200))
}
