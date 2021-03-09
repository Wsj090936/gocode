package functions

import (
	"fmt"
	"testing"
	"time"
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
	// 初始化结构体-引用
	person := &Person{
		age:  20,
		tall: 180,
	}
	// 初始化结构体-值
	//person := Person{
	//	age:  20,
	//	tall: 180,
	//}
	// 这里传入的是引用
	doPerson(person)
	person.getAge() // 这个人的年龄为： 20
	person.setAge(30)// 这里传入的是引用 函数内改变变量对外是有影响的
	person.getAge() // 这个人的年龄为： 30
}
func doPerson(p *Person) {
	p.age = 1
	fmt.Println(p.age)
}

type Person struct {
	age int
	tall int
}
// (p Person)表示当前这个方法是Person这个结构体里面的
func (p Person)getAge() {
	fmt.Println("这个人的年龄为：",p.age)
}
// 想要修改age 这里必须传指针,因为go语言默认的传递都是值传递，都是给函数一份拷贝
func (p *Person)setAge(age int)  {
	p.age = age
}
func TestPerson1(t *testing.T) {
	p := &Person{
		age:  10,
		tall: 20,
	}
	p.getAge()
	p.setAgeVal(11)
	p.getAge()
	num := 10
	incr(num)
	fmt.Println("num:",num)
	incrAge(*p)
	p.getAge()
}
func (p Person)setAgeVal(age int)  {
	p.age = age
}
func incr(num int) {
	num++
}
func incrAge(p Person) {
	p.age = 12
}



// 数组值传递测试
func TestArrRefrence(t *testing.T){
	arr := [...]int{1,2,3}
	transferArr(arr)
	// 结果还是 1,2,3
	fmt.Println(arr)
	transferArrF(&arr)
	fmt.Println(arr)// 4,2,3
	slice := make([]int,0,0)
	slice = append(slice,1,2,3)
	transferSlice(slice)
	fmt.Println(slice)
}
func transferSlice(slice []int) {
	slice[0] = -1
}

func transferArr(arr [3]int){
	arr[0] = 4
}
func transferArrF(arr *[3]int){
	arr[0] = 4
}

/**
	函数作为参数和返回值
 */

func TimeSpFunc(incr func(int) int) func(int) int {
	return func(n int) int {
		start := time.Now()
		res := incr(n)
		fmt.Println("当前函数运行时间:",time.Since(start).Seconds())
		return res
	}
}

func SlowFunc(num int) int {
	time.Sleep(time.Second * 2)
	fmt.Println("当前数字为:",num)
	return num
}

func TestTimeFunc(t *testing.T) {
	tsF := TimeSpFunc(SlowFunc)
	tsF(2)
}

func TestDefer(t *testing.T) {
	DeferFunc()
}

func DeferFunc() {
	defer SlowFunc(3)// 相当于finally
	fmt.Println("先执行")
	panic("error") //报错
}

func TestMultFunc(t *testing.T) {
	fmt.Println(Mulit())
}
func Mulit() (num int, tall int, err error) {
	num = 1
	tall = 2
	err = nil
	return
}



// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func TestSquares(t *testing.T) {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	f = squares()
	fmt.Println(f()) // "1"
}

func TestNiming(t *testing.T) {
	var x  = 2
	fun := func() int {
		x++
		return x * x
	}
	i := fun()
	fmt.Println(i)
	fmt.Println(x)
	i = fun()
	fmt.Println(i)
	fmt.Println(x)

}