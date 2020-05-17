package main

import (
	"fmt"
	someFunc "goStudy/importTest"
)

func main() {
	fmt.Println("Hello World")
	someFunc.SayHello()

	//const: 상수 var: 변수 let:?

	const name string = "Carlos Sainz"
	//name = "Charles Leclerc" 못바꿈
	fmt.Print(name)

	// name3 := "something -> var name3 string = "something"
	var name2 string = "Sainz"
	name3:= "Carlos"

	fmt.Println(name2)
	fmt.Println(name3)
}
