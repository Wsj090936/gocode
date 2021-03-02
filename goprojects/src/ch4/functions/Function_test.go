package functions

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	doRun(1)
	num1,num2 := getPeoples()
	fmt.Println(num1)
	fmt.Println(num2)
}
// 一个返回数
func doRun(pepole int) int {
	fmt.Println(pepole)
	return pepole + 1
}
// 两个返回值
func getPeoples() (int,int){
	return 1,2
}

// 测试引用传递
func TestRefrence(t *testing.T){
	a,b := 1,2
	fmt.Println("a和b未交换之前的值为:",a,b)
	// &代表取址符，取对应变量的地址
	swap(&a,&b)
	fmt.Println("a和b未交换之前的值为:",a,b)

}
func swap(a *int,b *int)  {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

/**
	函数作为参数传递
 */
type myFunc func(int) int
func TestFuncTransport(t *testing.T){
	funcVal := func(num int) int {
		return num * 2
	}
	fmt.Println(funcVal(2))
	doTrans(funcVal)
	doTransType(funcVal)
}

func doTrans(fun func(int) int){
	val := fun(3)
	fmt.Println("传递后的函数调用为:",val)
}
func doTransType(fun myFunc) {
	val := fun(4)
	fmt.Println("自定义type传递后的函数调用为:",val)
}

/**
	函数作为返回值
 */

func TestReturnFunc(t *testing.T){
	sumFunc := returnFunc()// 返回了一个函数
	fmt.Println(sumFunc(1,2))// 结果为3
	typeFunc := returnFuncType()
	fmt.Println(typeFunc(2,3))
}

type SumFunc func(int, int) int
func returnFunc() func(int,int) int{
	funcVal := func(a int,b int) int {
		return a + b
	}
	return funcVal
}
func returnFuncType() SumFunc{
	funcVal := func(a int,b int) int {
		return a + b
	}
	return funcVal
}
/**
	函数作为结构体的方法，将函数与结构体绑定了
 */

func TestPerson(t *testing.T) {
	person := &Person{
		age:  20,
		tall: 180,
	}
	person.getAge() // 这个人的年龄为： 20
}

type Person struct {
	age int
	tall int
}
// (p Person)表示当前这个方法是Person这个结构体里面的
func (p Person)getAge() {
	fmt.Println("这个人的年龄为：",p.age)
}
