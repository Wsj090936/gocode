package interfaces

// 实现排序接口

import (
	"fmt"
	"sort"
	"testing"
)

type MyIntList []int


func (p MyIntList)Len() int{
	return len(p)
}

func (p MyIntList)Less(i, j int) bool{
	return p[i] > p[j]
}
// Swap swaps the elements with indexes i and j.
func (p MyIntList)Swap(i, j int){
	p[i],p[j] = p[j],p[i]
}

func TestSelfSort(t *testing.T) {

	e := fmt.Errorf("xxx")
	list := MyIntList{}
	list = append(list,1)
	list = append(list,-1)
	list = append(list,3)
	list = append(list,2)
	list = append(list,5)
	sort.Sort(list)
	fmt.Println(list)
}
