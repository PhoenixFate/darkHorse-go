package main

import "fmt"

func main() {

	for i := 1; ; i++ {
		fmt.Println(i)
		if i > 20 {
			break
		} else {
			continue
		}
	}

	fmt.Println("21 break continue")
}
