package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("hello")
	//os.Exit(-1)
	if len(os.Args) > 1{
		fmt.Println(os.Args[1])
	}
}
