package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"reflect"
)

func main() {
	b := make([]byte, 10)
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

	hash := sha256.Sum256(b[:])
	//ed5cfbf9a26aad64ffcf76952c1a0232cfa206b1c779795c0788183815ad175e
	//对于任意长度的消息，SHA256都会产生一个256bit长的哈希值，称作消息摘要。
	//这个摘要相当于是个长度为32个字节的数组，通常用一个长度为64的十六进制字符串来表示
	fmt.Printf("hash: %x\n", hash)
	log.Panic("log panic test")
}
