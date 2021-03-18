package ch9

import (
	"fmt"
	"testing"
)

type X struct {
	A int
	B string
}
func BB(v X){
	fmt.Printf("%+v\n",v)
}
func AA()(v int){
	x := X{
		A: 10,
		B: "100",
	}

	x.A = 100
	x.B = "21312"
	defer BB(x)

	fmt.Printf("%+v\n",x)
	return v
}
func TestBB(t *testing.T) {
	AA()
}