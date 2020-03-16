package main

import "fmt"

func main() {
	//数组的元素个数必须是常量
	var id [20]int
	for i := 0; i < len(id); i++ {
		id[i] = i
	}
	for i := 0; i < len(id); i++ {
		fmt.Printf("id[%d]=%d\n", i, id[i])
	}

	const b = 10
	var arr [b]int
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]=%d\n", i, arr[i])
	}

	//定义数组的时候初始化
	//全部初始化
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d]=%d\n", i, arr2[i])
	}

	//部分初始化, 没有初始化的元素，默认为0
	arr3 := [5]int{1, 2, 3}
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr3[%d]=%d\n", i, arr3[i])
	}

	fmt.Println("04 array")
}
