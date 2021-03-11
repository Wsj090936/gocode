package interfaces

type MyApplePhone struct {
	*ApplePhone // 表示MyApplePhone继承了ApplePhone
	name string
}

func (myPhone MyApplePhone) getName() string {
	return myPhone.name
}
func (myPhone MyApplePhone) getNum() int {
	return myPhone.number * 10
}