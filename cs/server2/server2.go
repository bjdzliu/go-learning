package main

import (
	"net"
	"fmt"
	"strings"
)

func HandleConn(conn net.Conn){
defer conn.Close()
addr:=conn.RemoteAddr().String()
fmt.Println("addr conn successful",addr)
buf:=make([]byte,2048)
//循环读取
i:=0
//for循环的用途，一个client在不断开的情况下，一直使用这个连接
//client输入a回车。server端收到a和一个换行符号
for{
	//读取的内容，放入到buf中，返回长度
	fmt.Println("before conn.Read(buf)")
	n,err:=conn.Read(buf) //没有读取到内容时，阻塞

	if err!=nil{
		fmt.Println("no next input from client")
		fmt.Println("err=",err)
		return
	}
	fmt.Printf("read buf=%s\n",string(buf[:n])) //byte 转string类型
	if "exit"==string(buf[:n-2]){
		fmt.Println(addr,"exit")
		return
	}

	//将client传来的字符，转换成byte型，再传回给用户
	conn.Write([]byte(strings.ToUpper(string(buf[:n])))) 	//strings类型
i++
fmt.Println(i)
}

}

func main(){
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("err=",err)
		return
	}

	defer listener.Close()

	for{
		fmt.Println("before get listen")
		conn,err:=listener.Accept() //无输入时，阻塞等待状态
		fmt.Println("after listen")
		if err!=nil{
			fmt.Println("err=",err)
			return
		}
go HandleConn(conn)

	}




}
