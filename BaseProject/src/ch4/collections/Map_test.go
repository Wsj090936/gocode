package collections

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// map的初始化 map[键类型]值类型
	var m = make(map[string][]int)
	m["1"] = []int{1,2,3,4}
	m["2"] = []int{2,3,4}
	m["3"] = []int{3,4}
	fmt.Println(m)
	// 普通for循环的方式，idx表示键，e表示值
	for idx,e := range m{
		fmt.Println("键：" + idx + "值 ", e)
	}
	delete(m, "3")
	fmt.Println("键为3的值被删除")
	// 类似于while的方式 e表示键
	for e := range m {
		fmt.Println("键：" + e + "值 ",m[e])
	}
	// 这里相当于判断是否存在key为3的键，如果不存在，ok就返回false
	num,o := m["3"]
	fmt.Println(num)
	fmt.Println(o)

}