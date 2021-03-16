package main

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value int
}

func GenerateLinkedList() *ListNode {
	curr := &ListNode{Value: 0}
	front := curr
	for i := 1; i < 10; i++ {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Value = i
	}
	return front
}

func ReverseWithSlice(l *ListNode) *ListNode {
	var ptrslice []*ListNode
	for l != nil {
		ptrslice = append(ptrslice, l)
		l = l.Next
	}
	for i := len(ptrslice) - 1; i > 0; i-- {
		ptrslice[i].Next = ptrslice[i-1]
	}
	ptrslice[0].Next = nil
	return ptrslice[len(ptrslice)-1]
}

func ReverseLinkedList(l *ListNode) *ListNode {
	var prev, curr, next *ListNode
	curr = l
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}

func PrintLinkedListValues(l *ListNode) {
	for l != nil {
		fmt.Print(l.Value, "\t")
		l = l.Next
	}
	fmt.Println()
}

func main() {
	a := GenerateLinkedList()
	b := GenerateLinkedList()
	c := GenerateLinkedList()

	fmt.Println("Printing original:")
	PrintLinkedListValues(a)

	b = ReverseLinkedList(b)

	fmt.Println("Printing reversed:")
	PrintLinkedListValues(b)

	c = ReverseWithSlice(c)

	fmt.Println("Printing reversed with slice:")
	PrintLinkedListValues(c)
}
