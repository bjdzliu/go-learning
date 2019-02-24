package main

import (
	"time"
	"fmt"
)

/*
chan 在形参中
变量地址+channel
避免数据竞态
 */
func main(){
	cookedchan:=make(chan *Cake,1)  //这里指定只允许一个GR访问chan
	icechan:=make(chan *Cake,1)
	//cookedchan <- &Cake{"ok" }  这个不需要
	go baker(cookedchan)
	go icer(icechan,cookedchan)

	time.Sleep(20*time.Second)
}

type Cake struct{state string}
func baker(cooked chan<- *Cake){   //单向chan，写入。参数是一个chan
	for i:=0;i<10;i++{
		cake:=new(Cake)
		cake.state="cooked"
		cooked<-cake  //baker 不再访问cake变量;传递变量的地址给下一个GR
	}
	close(cooked)
}

func icer(iced chan *Cake,cooked <- chan *Cake){  //单向cooked chan，读取
	for cake:=range cooked{  //接收最后一个值后，关闭循环
		fmt.Printf("cake type is %T: \n",cake)
		cake.state="iced"
		fmt.Println(cake.state)
		iced<-cake
		//	iced <-  &Cake{"ok" }
		x,ok:=<- iced
		fmt.Println("after get iced",x.state,ok)
		//ok表示false，表示当前接收操作在一个关闭的并且读完的通道上。
		//为什么返回false? 这个GR已经执行put了，其它GR执行GET
		//icer先执行，等待baker，baker放入值，icer继续
		//baker放入值，等待icer的 range获取

	}
}