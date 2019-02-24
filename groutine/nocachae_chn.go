package main

import (
	"fmt"
	"time"
)

func main(){
	ch:=make(chan int,0)

	//fmt.Printf("%v,%v\n",len(ch),cap(ch))

	go func(){
		for i:=0;i<10;i++{
			fmt.Printf("in sub routine:%v\n",i) //1,8【in sub routine:1】,
			ch <- i //2(i=0,阻塞),9(i=1，阻塞),11(i=2，阻塞)，15,
			fmt.Printf("in sub routine len()=%v,cap()=%v\n",len(ch),cap(ch)) //7,10(得到时间切片)
		}
	}()
time.Sleep(2*time.Second) //3
	for i:=0;i<10;i++{
		num:=<-ch  //4（取0）,6（阻塞）,9.5(读到i=1，没有得到时间切片，还没来得及下一句。)12(num=1),14(阻塞)
		fmt.Printf("in main goroutine %v\n",num) //5（num=0）,13(num=1)

	}
}
