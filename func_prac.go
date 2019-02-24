package main

import (
	"fmt"
)

func main() {

	c(1, 2)
	fmt.Println(c)

}

//(int, string) 是返回值的类型
func A(a int, b string) (int, string) {

}

//abc是命名返回值参数
func B(a int, b string) (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

//不定长变参，放在最后一个参数位置,a是slice类型,a获得一个地址拷贝
//对a的修改不会影响到调用函数
//
func c(a ...int) (a, b, c int) {
	fmt.Println(a)
}

/////////////////////////////////
func main1() {
	s1 := []int{1, 2, 3, 4}
	A(s1)
	fmt.Println(s1)

}

//对slice s1地址进行拷贝
func a1(s []int) {
	s[0] = 5
	//修改了main1中的s1
	fmt.Println(s)
}

//////////////////////////////////
func main2() {
	a := 1
	//&a 取a的地址
	a2(&a)
}

func a2(a *int) {
	*a = 2 //修改了main2里面a的值
	fmt.Println(*a)
}

////////// 匿名函数 //////////////////
func mamain() {
	a := func() {
		fmt.Println("Func A")
	}
}

////// 闭包 /////
// 返回类型是 最右边的int
//定义一个函数closure,返回值是匿名函数func(int) int
//该匿名函数的返回值是int

func closure(x1 int) func(int) int {
	ffnei := 1
	return func(ffnei int) int {
		fmt.Println("Func A")
		return 1
	}
}
