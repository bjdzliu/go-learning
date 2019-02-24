package main

import (
	"fmt"
)

//为type类型进行方法绑定
// TZ 是int型的别名
type TZ int

type A struct {
	name string
}

func main() {
	var a TZ
	b := A{}

	a.Print()       //method value
	(*TZ).Print(&a) //method expression
	b.Print()
	fmt.Println(b.name)
}

//为int类型添加了print方法
func (a1 *TZ) Print() {

	fmt.Println("A")
}

//为struct添加方法,注意该方法可以访问小写字母name
func (b *A) Print() {
	b.name = "struct's name"
	fmt.Println(b.name)
}
