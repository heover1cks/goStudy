package main

import (
	"fmt"
	someFunc "goStudy/functionsTest/importTest"
	"strings"
)

func multiply(a int, b int) int {
	return a*b
}

func multiply2(a,b int) int { //input, return type 명시 필요
	return a*b
}

//Multiple Return value -> Go에만 있음
func lenAndUpper(name string)(int, string){
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) { //...
	fmt.Println(words)
}

func main() {
	fmt.Println("Hello World")
	someFunc.SayHello()

	//const: 상수 var: 변수 let:?

	const name string = "Carlos Sainz"
	//name = "Charles Leclerc" 못바꿈
	fmt.Println(name)

	// name3 := "something -> var name3 string = "something"
	var name2 string = "Sainz"
	name3:= "Carlos"

	fmt.Println(name2)
	fmt.Println(name3)

	//다중 리턴 펑션
	totalLength, upperName := lenAndUpper("Ricciardo")
	fmt.Println(totalLength,upperName)
	totalLength2,_ := lenAndUpper("Hulkenberg")
	fmt.Println(totalLength2)

	//repeatMe -> array type return
	repeatMe("Carlos", "Sainz","Ricciardo","Leclerc")
}
