package main

import (
	"fmt"
	"sync"
)

//同步包的形式，实现通信

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10) //创建10个任务
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait()
}

func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}
