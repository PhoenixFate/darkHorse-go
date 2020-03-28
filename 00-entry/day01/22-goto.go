package main

import "fmt"

func main() {

	fmt.Println("111111")
	fmt.Println("222222")

	//goto是关键字，End是用户起的名字
	goto End
	fmt.Println("can't print")

End:
	fmt.Println("3333")

	fmt.Println("22 goto")
}
