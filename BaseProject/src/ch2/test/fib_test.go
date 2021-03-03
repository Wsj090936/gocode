package test

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	//var a int = 1
	//var b int = 1

	//var (
	//	a int = 1
	//	b int = 1
	//)

	a,b := 1,1

	a := 1
	b := 1
	for i := 0; i < 5; i++ {
		fmt.Println(a)
		temp := a
		a = b
		b = temp + a
	}
}
