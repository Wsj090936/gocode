package interfaces

// 接口结构体，表示手机这一大类
type Phone interface {
	call() //
	getNum() int // 获取当前手机的手机号
}