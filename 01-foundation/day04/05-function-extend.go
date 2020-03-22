package main

import "fmt"

type Person4 struct {
	name string
	age  int
}

func (p *Person4) printInfo() {
	fmt.Println(*p)
}

//默认Student4 会继承Person4的方法
type Student4 struct {
	Person4
	id int
}

//方法的重写
func (s *Student4) printInfo() {
	fmt.Println(*s)
}

func main() {
	var p1 = Person4{"tom4", 88}
	p1.printInfo()

	var s1 = Student4{Person4{"abc", 99}, 123}
	//Student4继承了Person4的方法
	//显示调用
	s1.Person4.printInfo()
	s1.printInfo()

	//方法值：方法的入口地址，隐藏了接收者 函数指针指向的值
	pFunc := p1.printInfo //这个就是方法值 调用函数时 无需再传递接收者 隐藏了接收者
	pFunc()

	//方法表达式：方法的入口地址，但没有隐藏接收者
	pFunc2 := (*Person4).printInfo
	pFunc2(&p1)

	fmt.Println("05 方法的继承")
}
