package fortest

import (
	"fmt"
	"testing"
)

func TestMulCondition(t *testing.T){
	i := 3
	// go语言if预计支持条件
	if i = i + 3;i > 0{
		fmt.Println(i)
	}
}