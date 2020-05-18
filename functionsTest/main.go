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

//naked return -> 함수에 미리 선언
func nakedLenAndUpper(name string) (length int, uppercase string){
	//length := 1 -> 이미 Length가 생성되어 생성 불가
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func deferLenAndUpper(name string) (length int, uppercase string){
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
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

	//Naked Return
	fmt.Println(nakedLenAndUpper("Carlos Sainz"))

	//defer: func가 끝났을 때 할 행동 정의
	fmt.Println(deferLenAndUpper("Max Verstappen"))
}
