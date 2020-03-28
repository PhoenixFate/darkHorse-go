package main

import "fmt"

func main() {

	var a = 10
	var b = 30
	var c int
	var d int
	a = a + b
	b = a - b
	c = 2 * 3
	// 和C语言一样，整数
	d = 11 / 2
	e := 13 % 3
	var f = 1
	//只有后置的++ 后置的--
	f++
	f--
	fmt.Printf("a=%d; b=%d; c=%d; d=%d; e=%d; f=%d\n", a, b, c, d, e, f)

	b1 := a == b
	b2 := a != b
	b3 := a > b
	b4 := a >= b
	b5 := a < b
	b6 := a <= b
	fmt.Printf("b1:%v; b2:%v; b3:%v; b4:%v; b5:%v; b6:%v", b1, b2, b3, b4, b5, b6)

	//非0就是真
	c1 := b1 && b2
	c2 := b1 || b2
	c3 := !b1
	fmt.Printf("c1=%v; c2=%v; c3=%v", c1, c2, c3)

	fmt.Println("16")
}
