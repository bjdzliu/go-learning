package main

import (
	"time"
	"fmt"
)

/*
执行斐波那契数列，但是速度很慢
输出结果为 Fibonacci(45) = 1134903170
 */

func main() {
	go spinner(100 * time.Millisecond) //100毫秒

	fmt.Printf("\r%s", "asd")
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` { //输出过程，类似运行中状态
			fmt.Printf("\r%c", r) //%c 输出单个字符
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
