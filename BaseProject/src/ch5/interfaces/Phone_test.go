package interfaces

import (
	"fmt"
	"testing"
)

func TestPhone(t *testing.T) {
	apPhone := ApplePhone{number:10}
	myPhone := MyApplePhone{
		ApplePhone: &apPhone,// 表中MyApplePhone继承了ApplePhone
		name:       "我的手机",
	}
	apPhone.call()
	fmt.Println(apPhone.getNum())
	fmt.Println(myPhone.getNum())
}

func TestPhone2(t *testing.T) {
	var phone Phone
	phone = &ApplePhone{number: 2}
	s,ok := phone.(Phone)
	s.setNum(22)
	phone.setNum(22)
	fmt.Println(phone)

	fmt.Println(s,ok)
	fmt.Printf("%T",phone)
	p := ApplePhone{number:1}
	q := MyApplePhone{
		ApplePhone: &p,//
		name:       "我的手机",
	}
	q.number = 10 // 如果是引用，p中的信息也会发生变化
	fmt.Println(p)
	fmt.Println(q.number)
	num := q.getNum
	fmt.Println(num())
}



