package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	//使用a的方法
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println("b name is null", b.Name)
}

//方法Print和类型A绑定
//receiver is A
func (a1 *A) Print() {
	a1.Name = "xiu gai A Name"
	fmt.Println("A")
}

//b1.Name 值拷贝
//receiver is B
func (b1 B) Print() {
	b1.Name = "zhi chuandi"
	fmt.Println("B", b1.Name)
}
