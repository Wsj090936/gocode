package interfaces

import "fmt"

type ApplePhone struct {
	number int
}

func (applePhone ApplePhone) call() {
	fmt.Println("这是苹果手机")
}

func (applePhone ApplePhone) getNum() int {
	return applePhone.number
}
