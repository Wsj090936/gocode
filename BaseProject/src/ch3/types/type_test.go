package types

import (
	"fmt"
	"testing"
)

type myint int64

func Test_type(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	fmt.Println(b,a)
	var c string
	fmt.Print(len(c) == 0)
	// 指针
	aPtr := &a
	fmt.Printf("%T  %T",a,aPtr)
}