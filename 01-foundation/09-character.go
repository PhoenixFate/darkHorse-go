package main

import (
	"fmt"
)

func main() {
	var ch byte
	ch = 97

	fmt.Printf("character: %c %d\n", ch, ch)
	ch = 'b'
	fmt.Printf("character: %c %d\n", ch, ch)
	//大小写转换
	fmt.Printf("大写转小写: %c", 'A'+32)
}
