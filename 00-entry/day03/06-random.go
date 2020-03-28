package main

import "fmt"
import "math/rand"
import "time"

func main() {

	//设置种子
	//如果种子一样，每次的随机数一样
	//rand.Seed(10)
	//当前系统时间作为随机数
	rand.Seed(time.Now().UnixNano())

	//产生随机数
	fmt.Println("random int: ", rand.Int())
	//某个返回内的随机数
	fmt.Println(rand.Intn(100))

	fmt.Println("06 随机数")
}
