package operator

import (
	"fmt"
	"testing"
)

func Test_op(t *testing.T) {
	a := [...]int{1,2,3,4}
	b := [...]int{1,2,4,3}
	c := [...]int{1,2,3,4}
	//c := [...]int{1,2,4,3,5}
	fmt.Println(a == b)// false
	fmt.Println(a == c)// true
	//fmt.Println(c == b) // 这句不能运行
}
