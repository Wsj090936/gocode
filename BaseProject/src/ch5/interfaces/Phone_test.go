package interfaces

import (
	"fmt"
	"testing"
)

func TestPhone(t *testing.T) {
	apPhone := ApplePhone{number:10}
	myPhone := MyApplePhone{
		ApplePhone: apPhone,// 表中MyApplePhone继承了ApplePhone
		name:       "我的手机",
	}
	apPhone.call()
	fmt.Println(apPhone.getNum())
	fmt.Println(myPhone.getNum())
}