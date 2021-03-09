package oop

import (
	"fmt"
	"testing"
)

type Employee struct {
	Age int
	Name string
}

func TestEmployee(t *testing.T){
	em := new(Employee)

	fmt.Printf("type is %T",em)
}