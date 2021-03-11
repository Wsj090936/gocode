package structs

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	head := LinkedList{
		Val:  1,
		Next: nil,
	}
	head.Next = &LinkedList{2,nil}
	fmt.Println(head.getNextVal())
	next := head.getNext()
	next.Val = 3
	next.Next = &LinkedList{3,nil}
	fmt.Println(head.getNext())
}
