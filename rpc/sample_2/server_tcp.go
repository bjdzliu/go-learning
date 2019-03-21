package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/rpc"
	"net"
	"os"
)

func init(){
	fmt.Println("tcp server init")
}
type Data1 struct{
A,B int
}

type Data2 struct{
C,D int
}

type Num int

//传入对象 data1，返回整数
func (n *Num) Method1(args1 Data1,reply *int) error{
*reply=args1.A * args1.B
return nil
}

//传入对象 data1，返回 对象data2
func (n *Num) Method2(args1 Data1,reply *Data2) error{
	if args1.B==0{
		logs.Error("B err; is 0")
	}
	reply.C=args1.A/args1.B
	reply.D=args1.A%args1.B
	return nil
}

func main(){
	num:=new(Num)
	rpc.Register(num)
	tcpaddr,err:=net.ResolveTCPAddr("tcp",":9008")

	if err!=nil{
		fmt.Println("err occur")
		os.Exit(1)

	}
	listener,err:=net.ListenTCP("tcp",tcpaddr)
	for{
		conn,err:=listener.Accept()
		if err != nil {
			logs.Error("somethine wrong",err)
			continue
		}
		rpc.ServeConn(conn)
	}

}