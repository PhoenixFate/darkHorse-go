package main

//. 操作：这个包导入之后在调用的时候，可以省略包名
import . "fmt"

// 给包名起别名
import myTime "time"

func main() {

	Println("01")
	myTime.Sleep(1000)
	Println("11 point")

}
