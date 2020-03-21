package main

import "fmt"

//map传递的是引用
func myTest(m map[int]string) {
	delete(m, 2)
}

func main() {
	var map1 map[int]string
	fmt.Println(map1)
	fmt.Println("length map1: ", len(map1))

	//通过make创建map
	var map2 = make(map[string]string)
	map2["a"] = "aaa"
	map2["b"] = "bbb"
	map2["c"] = "ccc"
	fmt.Println(map2)
	fmt.Println("length map2: ", len(map2))

	//make创建map，指定长度（容量）
	//但不够用，也会自动扩容
	var map3 = make(map[string]string, 10)
	fmt.Println(map3)
	fmt.Println("length map3: ", len(map3))

	//定义的时候初始化
	var map4 = map[int]string{1: "abc", 2: "bcd", 3: "edf"}
	fmt.Println(map4)
	//修改某个key对应的值
	map4[1] = "ggg"
	fmt.Println(map4)

	//遍历map
	for key, value := range map4 {
		fmt.Printf("key: %d; value:%s\n", key, value)
	}

	//判断map中的key是否存在
	value, ok := map4[0]
	if ok {
		fmt.Println("key 存在, value: ", value)
	} else {
		fmt.Println("key 不存在")
	}

	//map删除某一个key
	delete(map4, 1)
	fmt.Println(map4)

	myTest(map4)
	fmt.Println(map4)

	fmt.Println("--- 15 map ---")
}
