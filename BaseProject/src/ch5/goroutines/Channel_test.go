package goroutines

import (
	"fmt"
	"testing"
)


func TestChannel(t *testing.T) {
	slice := []int{1,2,3,4,5,6}
	ch := make(chan int)// 创建一个channel
	go sum(slice[:len(slice)/2],ch)
	go sum(slice[len(slice)/2:],ch)
	x,y := <-ch,<-ch // 将channel接收到的值传给x,y两个变量
	fmt.Println(x,y)
}
func sum(arr []int, c chan int) {
	sum := 0
	for _, each := range arr{
		sum += each
	}
	c <- sum //向channel 发送消息
}
