package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	a,err := divide(1,0)
	fmt.Println(a)
	fmt.Println(err)
}
func divide(a, b int) (int, error) {
	if b == 0{
		return 0,errors.New("出错啦")
	}
	return a / b,nil
}