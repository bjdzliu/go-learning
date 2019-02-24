package main

import (
	"fmt"
	"time"
)

func newTask(){
	fmt.Println("this is new task")
}

func main(){
	go newTask()  //新建一个协程
	for {
		fmt.Println("this a main goroutin")
		//休眠 Duration的单位为 nanosecond。
		time.Sleep(2*time.Second)
		fmt.Println(time.Second)
		fmt.Println(time.Hour)
		fmt.Println("this a main goroutin")
		break

	}




}
