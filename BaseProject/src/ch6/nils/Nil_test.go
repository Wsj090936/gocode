package nils

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	var num *int
	fmt.Println(num)
	var num1 *int
	fmt.Println(num1)
	fmt.Println(gcd(24,60))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}