package duotai

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWord()
}

type GoProgrammer struct {

}

func (p GoProgrammer) WriteHelloWord()  {
	fmt.Println("go")
}

type JavaProgrammer struct {

}

func (p JavaProgrammer)WriteHelloWord()  {
	fmt.Println("java")
}

func TestProgrammer(t *testing.T) {
	goP := GoProgrammer{}
	javaP := JavaProgrammer{}
	Write(goP)
	Write(javaP)
}

func Write(p Programmer) {
	p.WriteHelloWord()
}