package main

import (
	"fmt"
	"time"
)

//var ch3 chan <- float64 只写
//var ch4 <- chan int 只读，只用于读取int型数据

var ch5 <- chan int

func counter(out chan <- int){  //只接收写channel
defer close(out)
for i:=0;i<10;i++{
	out<-i   //如果对方不读，会阻塞
	}
close(out)
}

//消费者，创建
func  read(in <- chan int){
	time.Sleep(time.Second)
for num:=range in{
	fmt.Println(num)
}


}

func main(){
	c:=make(chan int)
	//生产者：
	go counter(c)  //双向通道能转换程单向通道
	//消费者
	read(c)
	ch5 <- 6

}