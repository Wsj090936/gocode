package fortest

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	for i := 0;i < 7;i++{
		switch i{
		case 0, 1:
			fmt.Println("0,1")
		case 2, 3:
			fmt.Println("2,3")
		case 5, 4:
			fmt.Println("5,4")
		default:
			fmt.Println(i)
		}

	}
}