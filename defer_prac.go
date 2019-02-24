package main

import (
	"fmt"
)

////defer +  匿名函数 + 闭包
//// 结果输出 3 3 3
//func main() {

//	for i := 0; i < 3; i++ {
//		defer func() {
//在这个匿名函数里，并没有变量i的值，从外层获得。外层执行了三次循环，i变成3
//			fmt.Println(i)
//		}()
//	}
//}

func main() {
	// 非闭包，值拷贝
	for i := 0; i < 4; i++ {
		defer fmt.Println("value copy,i=", i)
	}

	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

func B() {
	// defer中 执行recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	// 遇到panic，程序中断，执行defer
	panic("panic in B")

}

func C() {
	fmt.Println("Func C")
}
