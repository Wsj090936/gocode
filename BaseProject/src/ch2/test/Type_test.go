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
