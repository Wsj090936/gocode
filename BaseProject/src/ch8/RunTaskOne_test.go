package ch8

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestTaskOne(t *testing.T) {
	fmt.Println("Before,",runtime.NumGoroutine())
	fmt.Println(AllResponse())
	time.Sleep(time.Millisecond*1000)
	fmt.Println("After,",runtime.NumGoroutine())

}
func runTask(i int) string {
	time.Sleep(time.Millisecond*50)
	return fmt.Sprint("The Complete task is ",i)
}
func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string,numOfRunner)// 如果不用buffer channel则一个被取走后，剩下的协程都会等待数据被取走
	for i := 0;i < numOfRunner;i++{
		go func(i int) {
			res := runTask(i)
			ch <- res
		}(i)
	}
	return <- ch
}


func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner) // 如果不用buffer channel则一个被取走后，剩下的协程都会等待数据被取走
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			res := runTask(i)
			ch <- res
		}(i)
	}

	res := ""
	for each := range ch {
		fmt.Println(each)
		res += "\n" + each
	}
	return res
}
