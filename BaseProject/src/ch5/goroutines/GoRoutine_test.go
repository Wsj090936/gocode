package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	go say("线程")// 开启一个goroutine 执行单独执行方法
	say("主线程")
}

func say(word string) {
	num := 5
	for num > 0 {
		num--
		time.Sleep(100 * time.Millisecond) // 每次睡100 ms
		fmt.Println(word)
	}

}
