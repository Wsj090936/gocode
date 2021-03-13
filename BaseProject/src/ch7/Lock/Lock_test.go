package Lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	var count  = 0
	for i := 0;i < 5000;i++{
		go func() {
			count++
		}()
	}
	time.Sleep(time.Second * 1)

	fmt.Println("非线程安全打印的结果为:",count)
	var countSafe = 0
	//var mut = sync.Mutex{}
	var mut sync.Mutex // 这句和上面那句等价 因为里面的参数被默认初始化了
	for i := 0;i < 5000;i++{
		go func() {
			defer func() {
				mut.Unlock()

			}()
			mut.Lock()
			countSafe++
		}()
	}
	time.Sleep(time.Second * 1)

	fmt.Println("线程安全打印的结果为:",countSafe)

}

func TestGroup(t *testing.T) {
	var countSafe = 0
	//var mut = sync.Mutex{}
	var mut sync.Mutex // 这句和上面那句等价 因为里面的参数被默认初始化了
	var group sync.WaitGroup
	for i := 0;i < 5000;i++{
		// 开启一个协程之前，增加一个等待量
		group.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			countSafe++
			group.Done() //执行完后，释放一个
		}()
	}
	group.Wait() // 注意 这里要等待所有线程完成
	fmt.Println("线程安全打印的结果为:",countSafe)
}
