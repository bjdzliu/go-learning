package main

import (
	"fmt"
	"net/rpc"
	"net/http"
	"errors"
)

func main(){
	rpcDemo()
}

type Arith int

//定义Arith的用途：

func rpcDemo(){
	arith:=new(Arith)
	fmt.Println("arith==",arith)

	rpc.Register(arith)
	rpc.HandleHTTP()
	err:=http.ListenAndServe(":9008",nil)
	if err!=nil{
		fmt.Println("err=====",err.Error())

	}
}

type Args struct{
A,B int
}
type Quotient struct{
	Quo,Rem int
}

func (t *Arith) Multiply(args *Args,reply *int) error{
	*reply = args.A * args.B
	fmt.Println("Multiply is executing",reply)
	return nil

}

func (t *Arith) Divide(args *Args,quo *Quotient) error {
	if args.B ==0 {
		return errors.New("Divided by zero")
	}
	quo.Quo = args.A/args.B //取整
	quo.Rem = args.A % args.B //求余
	return nil

}