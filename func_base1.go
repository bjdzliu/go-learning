package main

import (
	"fmt"
)

//必须有一个返回值
func test2(a int) int {
	//return fn()
	fmt.Println("sss", a)
	return 5
}

func test(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func main() {
	my := test2(8)
	fmt.Println(my)
	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprint(s, x, y)
	}, "%d,%d", 10, 20)

	fmt.Println(s1, s2)

	s3 := make([]int, 10, 20)
	s3 = []int{1, 2, 3}
	println(test3("sum: %d", s3...))
}

//变参
func test3(s string, gg ...int) string {
	var x int
	gg[3] = 100
	for _, i := range gg {
		x += i
	}
	return fmt.Sprintf(s, x)
}
