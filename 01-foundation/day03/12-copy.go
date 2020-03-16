package main

import "fmt"

func main() {

	s1 := []int{1, 2}
	s2 := []int{6, 6, 6, 6, 6}
	copy(s2, s1)
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("12 copy")
}
