package main

import (
	"fmt"
)

type Person struct {
	Name string
	Sex  byte
	Age  int
}

type master struct{
	Person
	Title string
	int
}



type Student struct {
	Person
	Name string
	id int
}

//
func (p Person) Setinfo(n string, s byte, a int) {
	p.Name = n
	p.Sex = s
	p.Age = a
}

// 接收者为指针变量,不能和上面的方法同名
func (p *Person) Setinfopoint(n string, s byte, a int) {
	p.Name = n
	p.Sex = s
	p.Age = a
}

func (p Person) Printinfo() {
	fmt.Println(p.Name, p.Age, p.Sex)
}

func main() {
	s1 := Person{"name", 'm', 22}
	stest:=Person{Name:"nan",Sex:'f',Age:33}
	fmt.Println(stest)

	fmt.Println(s1)
	(&s1).Setinfopoint("point", 9, 123)

	s2 := &Person{"methods", 'm', 100}
	//s2是指针，可以使用struct所有的方法，这些方法就是方法集
	//内部做了转换，先把指针p，转换成*p
	s2.Setinfo("newsetpoint", 'm', 20)
	fmt.Println("s2=", s2)
//匿名struct作为struct的成员在初始化时 需要带上struct的名字才行
	s3 := Student{Person:Person{"person'name",'m', 22}, id:100}
	s3.Printinfo()

	fmt.Println("s3.Person.Name",s3.Person.Name) //
	fmt.Println("s3.Name",s3.Name)


s3.Person.Name="person2's name"

	fmt.Println("after set s3.Person.Name",s3.Person.Name)
	fmt.Println("after set s3.Name",s3.Name)


	//master 匿名字段
	masterPerson:=master{Person{"",'1',10},"ii",100}
	fmt.Println(masterPerson)
//
	type A struct {a int}
	type B struct {a, b int}

	type C struct {A; B}
	var c C
	fmt.Println("c :",c.A.a)


}
