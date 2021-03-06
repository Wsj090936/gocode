package main
import (
	"fmt"
	"structs"
	//"ch5/interfaces"e
)

func main()  {
	fmt.Println("ok")
	p := structs.Person{}
	fmt.Println(p)
	var node *structs.TreeNode
	node = &structs.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(node.Val)
}