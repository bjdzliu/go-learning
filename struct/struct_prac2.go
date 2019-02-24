package main

import (
	"fmt"
	"strconv"
)

//结构体组合，类似继承
type human struct {
	Sex    int
	Name   string
	idcard int
	eyes string
}

type teacher struct {
	human //结构体的匿名字段
	int
	myint  // 自定义类型的匿名字段
	eyes string
	Age    int
	idcard int32
}

type student struct {
	human
	Name string
	Age  int
}
type myint int

func main() {
	a := teacher{ myint: 110, Age: 19, human: human{Sex: 1,Name:"human'name"}}

	b := student{Name: "hone", Age: 22, human: human{Sex: 22}}
	s0 := teacher{Age: 2}
	var s1 student = student{human: human{Sex: 22}, Name: "hone", Age: 22}
	var s2 student = student{human{Sex: 123}, "shunxu", 22}

	//s3 := student{human{Sex: 123}, "hone"}
	s3 := student{human{Sex: 123}, "hone", 22}

	fmt.Printf("s0=%+v\n", s0)
	fmt.Printf("s1=%+v\n", s1)
	fmt.Printf("s2=%+v\n", s2)
	fmt.Printf("s3= %v\n", s3)

	a.Name = "jone2" //teacher 本身没有Name，Name继承自human，所以这里修改了human的name
	fmt.Println("a.human.Name",a.human.Name)
	fmt.Println("a.teacher.Name",a.Name)

	a.eyes="teacher's eye"  //就近原则
	fmt.Println("a.eye.Name",a.eyes)   //只为teacher赋值
	fmt.Println("human.eye.Name",a.human.eyes)  //字符串为空


	a.Age = 56
	a.Sex = 100
	a.idcard = 6553611 //a.human.idcard = 999
	a.human.idcard = 1
	f := 100

	fmt.Println(a, b)
	fmt.Println(&a, &f)

	//打印出int的长度
	fmt.Println(strconv.IntSize)
}
