package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2) //o channel 保证程序不会迅速退出、看不到结果。
	go func() {
		for {
			select { //select放在无限循环里，不停地处理
			case v, ok := <-c1: //读取c1
				if !ok {
					fmt.Println("c1==") //close(c1)时，还执行了1-4次
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2: //读取c2
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)

			}

		}
	}()
	c2 <- "hi"
	c1 <- 1 // 写入chan，select会捕获到
	c1 <- 3
	c1 <- 4

	c2 <- "hello"

	close(c1) //关闭chan ， ok的值是false
	close(c2)
	for i := 0; i < 2; i++ {
		<-o
		fmt.Println("123")
	}

}
