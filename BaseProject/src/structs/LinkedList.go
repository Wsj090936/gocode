package structs

type LinkedList struct {
	Val int
	Next *LinkedList
}

func (p *LinkedList) getNextVal() int {
	return p.Next.Val
}

func (p *LinkedList) getNext() *LinkedList {
	return p.Next
}