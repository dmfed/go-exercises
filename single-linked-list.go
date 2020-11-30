package main

import (
	"errors"
	"fmt"
)

// ErrorEmpty is returned when PopFront is called on empty list
var ErrorEmpty = errors.New("list is empty")

// LinkedList is a plain implementation of single linked list with pushback and
// popfront methods
type LinkedList struct {
	front  *LinkedListElement
	back   *LinkedListElement
	length int
}

// LinkedListElement represents element of LinkedList structure
type LinkedListElement struct {
	next  *LinkedListElement
	value interface{}
}

// PushBack appends value to the end of the list
func (l *LinkedList) PushBack(payload interface{}) error {
	elem := LinkedListElement{next: nil, value: payload}
	if l.front == nil {
		l.front = &elem
	}
	if l.back != nil {
		l.back.next = &elem
	}
	l.back = &elem
	l.length++
	return nil
}

// PopFront extracts head of Linked List
func (l *LinkedList) PopFront() (interface{}, error) {
	if l.front == nil {
		return nil, ErrorEmpty
	}
	payload := l.front.value
	l.front = l.front.next
	l.length--
	return payload, nil
}

// Len returns number of elements surrently in the list
func (l *LinkedList) Len() int {
	return l.length
}

func main() {
	var l LinkedList
	for i := 0; i < 6; i++ {
		fmt.Println(l.PushBack(i))
	}
	fmt.Println(l.Len())
	for i := 0; i < 7; i++ {
		fmt.Println(l.PopFront())
	}
}
