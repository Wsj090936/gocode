package test

import (
	"fmt"
	"testing"
)

const(
	ONE = iota
	TWO
	THREE
)
const(
	OPEN = 1 << iota // 表示之后连续的常量都左移一位
	CLOSE
	RUNNING
)
func TestConst(t *testing.T) {
	fmt.Println(ONE,TWO,THREE)
	fmt.Print(OPEN,CLOSE,RUNNING)
}