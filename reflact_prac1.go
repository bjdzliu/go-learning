package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello")

}

func main() {
	var a User
	a.Age = 100
	fmt.Println(a.Age)
	Info(a)

	//var u = User{1, "ok", 12}

}

//根据变量a，获得a的类型，类型名
//o可以是任何struct
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type Name:", t.Name())         //User
	fmt.Println("Type String:", t.String())     //main.User
	fmt.Println("Type NumFiled:", t.NumField()) //3
	//Type.Field()返回的是StructFiled对象，有Name,Type等属性
	fmt.Println("Type Filed:", t.Field(1)) //{Name  string  8 [1] false}
	//判断接收的参数时候是struct
	if k := t.Kind(); k != reflect.Struct {

	}
	//ValueOf返回一个 value 类型
	v := reflect.ValueOf(o)
	fmt.Println(v) //{0  1}
	fmt.Println("Field below:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		//v.Field(i)返回的还是一个reflect.value类型，第i个field 。
		//调用Interface()方法，再反射回interface{}类型
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val) //Age int 100
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		//该方法对应的函数中，第一个接收的参数就是User
		fmt.Println(m.Name, m.Type) //Hello func(main.User)
	}

}
