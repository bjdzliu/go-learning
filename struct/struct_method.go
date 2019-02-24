package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

//type 	pointer *int
//func (tmp pointer) test(){}

func (tmp Person) Printfo(a int) {
	fmt.Println("tmp=", tmp)
}

//func (tmp Person) Printfo() {
//	fmt.Println("tmp=", tmp)
//}

func (tmp *Person) Setinfo(name string, sex byte, age int) {
	tmp.age = age
	tmp.name = name
	tmp.sex = sex

}

func main() {
	p := Person{"kk", 'm', 100}
	p.Printfo(1)
	p.name = "new"
	p.Printfo(1)
	var p2 Person
	(&p2).Setinfo("second", 'f', 99)
	p2.Printfo(1)

}
