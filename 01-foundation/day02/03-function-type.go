package main

import "fmt"

func MyAdd(a int, b int) int {
	return a + b
}

func MyMinus(a int, b int) int {
	return a - b
}

//MyFunction是一个函数类型
type MyFunction func(int, int) int

func main() {

	//声明一个函数类型的变量
	var funcTest MyFunction
	funcTest = MyAdd
	result := funcTest(10, 20)
	fmt.Println("result: ", result)

	fmt.Println("03 function type")
}
