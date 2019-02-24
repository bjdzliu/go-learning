package main

import (
	"net"
	"fmt"
	"os"
)

func main(){
//主动连接服务器
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("err=",err)
		return
	}
	//main调用完毕 ，关闭连接
defer conn.Close()
	//开启一个协程进行数据的获得
	buf:=make([]byte,1024)


	//从键盘输入，给服务器发送
go func(){

str:=make([]byte,2048)
for{
	//fmt.Println("str is",str)
	n,err:=os.Stdin.Read(str) //阻塞等待
	fmt.Println("running in routin",n)
	if err!=nil{
		fmt.Println("err=",err)
		return
	}
	//给server发送
	conn.Write(str[:n])
}
}()

	//问题：为什么输入和获取不能交换协程

	//
	for{
		n,err:=conn.Read(buf) //阻塞等待

		if err!=nil{
			fmt.Println("err=",err)
			return
		}
		fmt.Println("conn recv is ",string(buf[:n]))
	}

}
