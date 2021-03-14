package ch8

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("创建一个新的对象")
			return 3
		},
	}
	pool.Put(2)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())

}
