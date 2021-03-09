package oop

import (
	"fmt"
	"testing"
)

//type Programmer interface {
//	helloWord() string
//}

type JavaProgrammer struct {

}
type GoProgrammer struct {
	JavaProgrammer
}

func (p *JavaProgrammer) helloWord() string {
	fmt.Println("word")
	return "fmt.Printf(\"hello\")"
}

func (p *JavaProgrammer) Word(name string) string {
	p.helloWord()
	fmt.Println(name)
	return "fmt.Printf(\"hello\")"
}

func (p *GoProgrammer) helloWord() string {
	fmt.Println("go word")
	return "fmt.Printf(\"hello\")"
}
func TestClient(t *testing.T)  {
	//var p Programmer = new(JavaProgrammer)
	//p = new(JavaProgrammer)
	//fmt.Println(p.helloWord())
}
// 仅仅是类似继承 不支持LSP
func TestExtends(t *testing.T) {
	p := GoProgrammer{}
	p.Word("xx")// 这里其实调用 helloword 调用的还是javaProgrammer的方法
}