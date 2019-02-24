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

type Manager struct {
	User  //匿名字段，名字和类型都是User
	title string
}

func main() {
	m := Manager{User: User{3, "myname", 99}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Println(t.Field(0)) //{User  main.User  0 [0] true}  true表示是匿名字段
	//想办法取得user的id,用slice
	fmt.Println(t.FieldByIndex([]int{0, 0}))
	//想办法通过reflect修改内容
	x := 123
	v := reflect.ValueOf(&x) //传指针
	v.Elem().SetInt(999)
	fmt.Println(x)
}
