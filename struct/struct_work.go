// 如果匿名字段和外层结构有同名字段，应该如何操作
package main

import (
	"fmt"
)

//A的级别>B1>C
type A struct {
	B1
	//如果也放入C，那么有重名的错误
	//Name int  如果写这个，就找到这个
}

type B1 struct {
	C
	Name string
}

type C struct {
	Name string
}

func main() {

	a := A{B1: B1{Name: "B"}}
	b := B1{C: C{Name: "cc"}}

	//a.Name 在A的级别上找Name，如果找不到Name，在B1里找Name
	fmt.Println("a.Name =", a.Name)

	fmt.Println("a.B1.Name:", a.B1.Name)
	// b.Name 是str
	if b.Name == "" {
		fmt.Printf(" before set B's name =%q\n", b.Name)
		b.Name = "define b name"
		fmt.Printf(" after set B's name =%q\n", b.Name)
	}

	fmt.Println(b.C.Name)
}
