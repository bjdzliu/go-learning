package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

//func main() {
//	//初始化,有逗号
//	a := person{
//		Name: "joe",
//		Age:  12,
//	}
//	a.Age = 1
//	a.Name = "zhao"
//	fmt.Println(a)
//	A(a)
//	fmt.Println(a)
//}

////值拷贝
//func A(per person) {
//	per.Age = 13
//	fmt.Println("A person", per)
//}

var a1=1

// a2:=3

func main() {
	//初始化,有逗号，加&，a是地址
	//加&是为了下面的A(a)实现引用类型的copy
	//初始化方式一：
	a := &person{
		Name: "joe",
		Age:  12,
	}

	//*a.Age 简化a.Age = 1也是可以的
	(*a).Age = 1
	a.Name = "zhao"
	fmt.Println("a is ",a)
	A1(a)
	fmt.Println(a)


	//初始化方式二：
	init2:=person{
		Name:"mike",
		Age:12,
	}
	init2.Name="mike2"  //
	init2.Age=100
fmt.Println("init2:",init2)

	//方式三：
	ms := new(person)
	ms.Age=100
	(*ms).Name="ms name"
	fmt.Println(ms)

	B()
	// a已经是指针类型，不能在使用 a:= person3{}定义
	b := person3{"jone", 22}
	fmt.Println(b)
fmt.Println(a)
	var a3=100
	fmt.Println(a3)
}

//地址拷贝
func A1(per *person) {
	per.Age = 13
	fmt.Println("A person", per)
}

type cc struct{
	Name string
	Age int
}
var a=&struct{
Name string
Age int
}{
Name:"jone",
Age:19,
// 逗号不能省略
}
//匿名结构的用途

func B() (*cc) {
	//匿名结构

	//函数中可以用var定义
	//a是指针
	var a=&struct{
		Name string
		Age int
	}{
		Name:"jone",
		Age:19,
		// 逗号不能省略
	}

cc1:=&cc{"cc1",100}
fmt.Printf("B():%v   %v\n",a.Name,a.Age)
fmt.Printf("%v\n",cc1)
//return a
return cc1

}

//匿名结构嵌套，Contact的类型是匿名结构
func C1() {
	type person1 struct {
		Name string
		Age  int
		Contact struct {
			Phone, City string
		}
	}
	a := person1{Name: "jj", Age: 11}

	// 对匿名结构的初始化
	a.Contact.Phone = "123123"
	a.Contact.City = "beijing"

}
	//匿名字段,赋值时看顺序
	type person3 struct {
		string
		int
	}







