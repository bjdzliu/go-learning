package main

import "fmt"

type mystr string

type Student struct {
	//匿名字段 是struct类型指针
	*Person
	id   int
	addr string
	age string   //同名
}

type Person struct {
	name string
	sex  byte
	age  int  //同名
}

func main() {
	s := Student{&Person{"person1", 1, -2}, 22, "student addr","123"}
	//s.age 输出的是Student中的值
	fmt.Println(s.name, s.sex, s.age,s.addr,s.age)

	s.Person = &Person{"go2", 'm', 44}
	fmt.Println(s.name, s.sex, s.age,s.Person.age)

	s1:=Student{Person:&Person{"person1",1,2},addr:"addr2name",age:"s3's age"}
	fmt.Println(s1.Person,s1.age)
}
