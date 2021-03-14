package singleton

import (
	"fmt"
	"sync"
)

var once  = sync.Once{}
var singleTon *SingleTon
type SingleTon struct {

}


func GetInstance() *SingleTon {
	once.Do(func() {
		fmt.Println("单例创建")
		singleTon = new(SingleTon)
	})
	return singleTon
}
