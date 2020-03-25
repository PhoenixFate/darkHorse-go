package main

import "fmt"

func testA() {
	fmt.Println("aaa")
}

func testB(index int) {
	fmt.Println("bbb")
	//显示调用panic函数，会中断程序
	//panic("中断程序 after testB")

	//隐士调用
	var a [10]int
	fmt.Println(a[index])
}

func testC() {
	fmt.Println("ccc")
}

func main() {
	//panic: 严重的错误，一般当panic异常发生时，程序会中断
	testA()
	testB(20)
	testC()

	fmt.Println("02 panic")
}
