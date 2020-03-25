package main

import "fmt"

func testA2() {
	fmt.Println("aaa")
}

func testB2(index int) {
	//设置recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("bbb")
	//显示调用panic函数，会中断程序
	//panic("中断程序 after testB")
	//隐士调用
	var a [10]int
	fmt.Println(a[index])
}

func testC2() {
	fmt.Println("ccc")
}

func main() {
	//使用panic 会导致程序中断
	//使用recover继续程序
	testA2()
	testB2(29)
	testC2()

	fmt.Println("03 recover")
}
