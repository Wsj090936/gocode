package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := AsyncServide()
	OtherTask()
	fmt.Println( <- ch)// 这里也会进行等待
	fmt.Println("主线程执行完毕")
}

func service() string  {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func AsyncServide() chan string {
	var ch = make(chan string,1)
	go func() {
		res := service()
		fmt.Println("Async Service start")
		ch <- res  // 如果channel没有buffer，这里会一直阻塞直到被取走
		fmt.Println("Async Service end")
	}()
	return ch
}

func OtherTask() {
	fmt.Println("Other task start")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Other task complete")
}
