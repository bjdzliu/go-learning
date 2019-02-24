//通过reflect调用方法

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

func (u User) Hello(name string) {
	fmt.Println("hello", name, "my name is ", u.Name)
}

type Manager struct {
	User  //匿名字段，名字和类型都是User
	title string
}

func main() {
	u := User{1, "ok", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	// 参数要求传入一个slice，slice元素的类型是reflect.value
	args := []reflect.Value{reflect.ValueOf("join")}
	mv.Call(args)

}
