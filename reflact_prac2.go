package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	X string
	Y int
}

func main() {
	var i int = 123
	var ptr = &i
	var f float32 = 1.23
	var l []string = []string{"a", "b", "c"}

	fmt.Println(reflect.TypeOf(i))   //int
	fmt.Println(reflect.TypeOf(ptr)) //*int
	fmt.Println(reflect.ValueOf(i))  //123
	fmt.Println(reflect.TypeOf(f))   //float32
	fmt.Println(reflect.TypeOf(l))   //[]string

	var foo Foo
	foo2 := Foo{}
	foo3 := Foo{"x", 1}
	fmt.Println(reflect.TypeOf(foo2))       //main.Foo
	fmt.Println(reflect.TypeOf(foo).Name()) //Foo 返回结构体的名字
	fmt.Println(foo)                        //{ 0}
	fmt.Println(reflect.ValueOf(&foo3).Elem())

}
