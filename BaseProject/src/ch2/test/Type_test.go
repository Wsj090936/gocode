package test

import (
	"fmt"
	"structs"
	"testing"
)

func TestType(t *testing.T) {
	var person structs.Person
	person =  structs.GetNewPerson()
	fmt.Println(person)
	person.SetAge(2)
	fmt.Println(person)


}

func TestString(t *testing.T){
	var str = "世界"
	for i,r := range str {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	var slice []int = []int{}
	slice = append(slice,1)
	slice = append(slice,1)
	slice = append(slice,1)
	slice = append(slice,1)
	slice = append(slice,1)
	fmt.Println(len(slice),cap(slice))
	fmt.Println(slice[:10])
}
