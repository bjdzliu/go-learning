package main

import (
	"fmt"
	"reflect"
)

//struct类型，通过ValueOf

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
	var p *int
	var i int

	p = &i //指针p存放i的地址

	fmt.Printf("p: %v,%v\n", p, *p)
	//使用序号找到匿名字段
}

//处理struct类型，通过reflect

func Set(o interface{}) {
	vv := reflect.ValueOf(o)
	//fmt.Println(vv.NumField()) //panic: reflect: call of reflect.Value.NumField on ptr Value
	vv = vv.Elem()
	fmt.Println(vv.NumField())
	//vv = vv.Elem().NumField()   报错

	v := reflect.ValueOf(o)        //获取对象o的value，类型是value。v是value类型的指针，才可以调用Elem。
	fmt.Println(v)                 //&{1 OK 12}
	fmt.Println("ceshi", v.Elem()) //ceshi {1 OK 12}
	fmt.Println(v.Kind())          //ptr
	fmt.Println(v)                 //&{1 OK 12}
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return

	} else {
		v = v.Elem() //通过Elem才可以赋值
	}
	f := v.FieldByName("Name")
	if !f.IsValid() { //如果没有找到Name
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String {
		fmt.Println(v.NumField()) //struct属性的数量
		f.SetString("byebye")

	}
}
