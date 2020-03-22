package main

import "fmt"

//面向过程
func Add01(a, b int) int {
	return a + b
}

type long int

//面向对象的方法
//给某个类型绑定函数
func (temp long) Add02(b int) int {
	return int(temp) + b
}

func main() {
	var result int
	result = Add01(23, 23) //普通的函数调用方式
	fmt.Println(result)

	var a long = 10
	print(a.Add02(20))
	fmt.Println("03 对象中的函数")
}
