package errors

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	msg := zeroFunc()
	fmt.Println("msg的值为:",msg)
}
func zeroFunc() string {
	msg := "默认"
	fun := func() string {
		err := recover()
		if err != nil {
			fmt.Println(err)
			msg = "出现问题了"
		}
		msg = "问题不大"
		fmt.Println("匿名函数修改:", msg)
		return msg
	}
	defer fun()
	//byZero(0)

	return msg
}
//func byZero(num int) {
//	i := 1
//	i2 := i / num
//	fmt.Println(i2)
//}

