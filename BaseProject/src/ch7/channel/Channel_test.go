package channel

import (
	"fmt"
	"sync"
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


func TestSelect(t *testing.T){
	// 以下select 语句 case1 和case 2 会随机执行一个
	select {
	case ret := <- AsyncServide():
		fmt.Println("执行了1:",ret)
	case ret2 := <- AsyncServide():
		fmt.Println("执行了2",ret2)
	// 表示如果其他case超时了，就执行这条语句
	case <-time.After(time.Millisecond * 100):

		fmt.Println("失败了，执行后面吧")
	// default 表示如果所有case都没准备好，就执行default
	default:
		fmt.Println("我是default")
	}
	fmt.Println("继续执行后面")
}

func TestChannelClose(t *testing.T) {
	var ch = make(chan int)
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go doProducer(&wg,ch)
	wg.Add(1)
	go doReceiver(&wg,ch)
	wg.Wait()
}
func doProducer(wg *sync.WaitGroup,ch chan int) {
	for i:=0;i < 10;i++{
		ch <- i
	}
	// 这里进行关闭通道
	close(ch)
	wg.Done()
}

func doReceiver(wg *sync.WaitGroup,ch chan int) {
	for{
		// 这里如果关闭了，就会进行关闭的广播
		if data,ok := <- ch;ok{
			fmt.Println(data)
		}else {
			fmt.Println("程序退出了")
			break
		}
	}
	wg.Done()
}