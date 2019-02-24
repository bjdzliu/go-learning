package main

import (
	"time"
	"fmt"
)

func main(){
	//2秒后，向time通道写内容
	timer1:=time.NewTimer(2*time.Second)
	t0:=time.Now()
	fmt.Printf("t1:%v\n",t0)

	//2s后，执行这句话
	t1:=<-timer1.C  //没有数据时，阻塞
	fmt.Println(t1)

	//延迟功能：
	<-time.After(2*time.Second)
   //2s后，产生一个事件
    fmt.Println("时间到")

}
