package functions

import (
	"fmt"
	"testing"
)

func TestDeger(t *testing.T) {
	defer fmt.Println("退出前执行")

	defer func() {
		fmt.Println("执行了")
	}()

	fmt.Println("退出了")
}
