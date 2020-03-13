package main

import "fmt"

func main() {
	fmt.Println("08")

	var f1 float32
	f1 = 3.24
	fmt.Println("f1= ", f1)

	//自动类型推导，浮点型的默认float64
	f2 := 5.23
	fmt.Printf("f2 type:%T\n", f2)

}
