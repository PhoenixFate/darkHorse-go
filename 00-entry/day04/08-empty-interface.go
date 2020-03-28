package main

import (
	"fmt"
	"reflect"
)

type Teacher8 struct {
	group string
}

//空接口 interface{} 类似于c语言中的void*
//任何类型都实现了空接口，因此空接口可以存储任意类型

func MyPrint(args ...interface{}) {
	fmt.Println(args)
}

func main() {
	//空接口 可以保存任意类型
	var v1 interface{} = 1
	fmt.Println(v1)
	var v2 interface{} = "abc"
	fmt.Println(v2)

	MyPrint(v1)
	MyPrint(v2)
	MyPrint(v2, "ddd")

	//类型查询
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "abc"
	i[2] = Teacher8{"sss"}
	//类型判断
	for index, data := range i {
		fmt.Println(reflect.TypeOf(data))

		if value, ok := data.(int); ok == true {
			fmt.Printf("i[%d] is int; value=%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("i[%d] is string; value=%v\n", index, value)
		} else if value, ok := data.(Teacher8); ok == true {
			fmt.Printf("i[%d] is Teacher8; value=%v\n", index, value)
		}

	}

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("i[%d] is int; value=%d\n", index, value)
		case string:
			fmt.Printf("i[%d] is string; value=%v\n", index, value)
		case Teacher8:
			fmt.Printf("i[%d] is Teacher8; value=%v\n", index, value)
		}
	}

	fmt.Println("08 空接口")
}
