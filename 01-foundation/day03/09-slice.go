package main

import "fmt"

func main() {

	//[low:high:max]
	//low 下标起点
	//high 下标终点（不包括）
	//len=high-low 长度
	//cap容量 max-low

	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:3:5]
	fmt.Println("s= ", s)
	fmt.Println("len: ", len(s))
	fmt.Println("cap: ", cap(s))

	s = a[1:4:5]
	fmt.Println("s= ", s)
	fmt.Println("len: ", len(s))
	fmt.Println("cap: ", cap(s))

	//数组和切片的区别
	//数组不能改变长度 []里面必须是固定值
	b := [5]int{}
	fmt.Println("len b: ", len(b))
	fmt.Println("cap b: ", cap(b))

	//切片 []里面为空，或者...
	var s2 []int
	fmt.Println("len s2: ", len(s2))
	fmt.Println("cap s2: ", cap(s2))
	s2 = append(s2, 22)
	fmt.Println("len s2: ", len(s2))
	fmt.Println("cap s2: ", cap(s2))

	//切片的创建
	//自动推导类型并且初始化
	s3 := []int{12, 3, 4, 5}
	fmt.Println(s3)
	//借助make函数 make(切片类型, 长度, 容量）
	s4 := make([]int, 5, 10)
	fmt.Println("len s4: ", len(s4))
	fmt.Println("cap s4: ", cap(s4))

	//切片的操作
	//1.截取
	//s[n] 索引为n的项
	//s[:] 0 ~ len-1
	//s[low:] low ～ len-1
	//s[:high] 0 ~ high
	//s[low:high] low:high
	//s[low:high:max]

	myArray := []int{1, 2, 3, 4, 5, 6, 7}
	s5 := myArray[:]
	fmt.Printf("len=%d, cap=%d, s5:%v\n", len(s5), cap(s5), s5)
	s6 := myArray[0]
	fmt.Printf("s6:%v\n", s6)
	s7 := myArray[3:6:7]
	fmt.Printf("len=%d, cap=%d, s7:%v\n", len(s7), cap(s7), s7)

	s8 := myArray[:4]
	fmt.Printf("len=%d, cap=%d, s7:%v\n", len(s8), cap(s8), s8)

	fmt.Println("09 slice")
}
