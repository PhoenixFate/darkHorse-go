package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := make([]byte, 100)
	b[0] = 'a'
	b[1] = 'c'
	b[2] = 'm'
	a := 10
	c := &a
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(b[:]))

	fmt.Println(string(b))
	fmt.Println(string(b[:]))
	fmt.Println(string(append(b, 'i', 'l', 'v', 'u')))
	fmt.Println(string(b[:]))
}
