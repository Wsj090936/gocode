package ch9

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectTest(t *testing.T) {

	var num = 10
	of := reflect.ValueOf(num)
	fmt.Println(of.Int())
	e := &Employee{
		Age:  18,
		Name: "mike",
	}
	field := reflect.ValueOf(*e).FieldByName("Age")// 获取值必须要用 实体对象
	field2 := reflect.ValueOf(e).Elem().FieldByName("Age").Int()
	fmt.Println("通过反射获取值",field2)
	// 设置值必须要用指针对象并取出其地址
	age := reflect.ValueOf(e).Elem().FieldByName("Age")
	age.SetInt(3)
	fmt.Println("反射设置值之后:{} ",*e)
	fmt.Println("反射获取值：",field.Int() )
	method := reflect.ValueOf(e).MethodByName("UpdateAge")// 这里设置值一定要传入引用
	method.Call([]reflect.Value{reflect.ValueOf(2)})
	// 第二种设置值的方式
	fmt.Println(*e)


}

type Employee struct {
	Age int
	Name string
}

func (e *Employee) UpdateAge(age int)  {
	e.Age = age
}
