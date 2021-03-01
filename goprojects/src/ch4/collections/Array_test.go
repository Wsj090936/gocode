package collections

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	var arr [3]int
	fmt.Println(arr[0],arr[1])
	// 申明时初始化值
	arr1 := [4]int{1,2,3,4}
	fmt.Println(arr1[0],arr1[1])
	// 申明时指定随意的长度
	arr2 := [...]int{1,2,3,4,5}
	fmt.Println(arr2[4],arr2)
}

// 数组元素的遍历
func TestArrayTravel(t *testing.T)  {
	arr2 := [...]int{1,2,3,4,5}
	// 传统的遍历方式
	for i := 0;i < len(arr2);i++{
		fmt.Println(arr2[i])
	}
	// for each 遍历数组，idx表示当前值的下标，e表示当前值
	for idx,e := range arr2{
		fmt.Println(idx,e)
	}
}
// 数组的截断操作
func TestArraySection(t *testing.T) {
	arr2 := [...]int{1,2,3,4,5}
	// 表示从下标为0的元素开始，截取到下标为2的元素(不包括下标为3的)
	fmt.Println(arr2[:3])
	// 表示从下标为1的元素开始，截取到下标为2的元素(不包括下标为3的)
	fmt.Println(arr2[1:3])
	// 表示从下标为1的元素开始，截取到最后一个元素
	fmt.Println(arr2[1:])

	a2 := arr2[1:]
	a2[2] = 0
	fmt.Println(arr2,a2) // [1 2 3 0 5] [2 3 0 5]
}
