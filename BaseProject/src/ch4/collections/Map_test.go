package collections

import (
	"encoding/json"
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

func TestMapExist(t *testing.T){
	m := map[int]int{}
	m[0] = 0
	if v,ok := m[1];ok{
		fmt.Println("Key  exist,value = ",v)
	}else {
		fmt.Println("Key not exist")
	}

}

func TestMySet(t *testing.T) {
	set := Set{m: map[int]bool{}}
	fmt.Println("检查 key 为 1 存不存在:",set.contains(1))
	set.add(1)
	fmt.Println("检查 key 为 1 存不存在:",set.contains(1))
	set.delete(1)
	fmt.Println("检查 key 为 1 存不存在:",set.contains(1))
	set.setName("吴世佳")
	fmt.Println(set.name)

}

func TestMapRefrence(t *testing.T)  {
	m := map[int]int{}
	fmt.Println(m)
	setMap(m,1,1)
	fmt.Println(m)
	bytes, e := json.Marshal(m)
	if e != nil{
		return
	}

	fmt.Println(string(bytes))
}
// map实际上也是值传递，但是里面有个值是指针。。所以有点像引用
func setMap(m map[int]int,key int,value int)  {
	m[key] = value
}