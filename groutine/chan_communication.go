package main

import (
	"fmt"
	"time"
)

func main(){
//创建channel,无缓冲区的，同步，读和写互相等待
ch:=make(chan string)
defer fmt.Println("main go routin end")
	go func(){
	for i:=0;i<2;i++{
		fmt.Println("i=",i)
		time.Sleep(time.Second)
		}
		ch <- "end"
		fmt.Println("in sub")
		//ch <- "end2"
	}()
//str:=<-ch
time.Sleep(3*time.Second)
    str2:=<-ch
	fmt.Println("get the str value is:",str2)

    fmt.Println(str2)
}
