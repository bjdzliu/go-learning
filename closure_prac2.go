package main

import (
	"fmt"
)

func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(int) int {
	fmt.Println(&x)
	return func(y int) int {
		fmt.Println(&x)
		return x + y
	}
}
