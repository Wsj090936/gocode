package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func TestSinle(t *testing.T) {
	var wg = sync.WaitGroup{}
	for i := 0;i < 10;i++{
		wg.Add(1)
		go func() {
			obj := GetInstance()
			fmt.Println(unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()

}