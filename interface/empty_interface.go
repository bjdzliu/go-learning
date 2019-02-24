package main

import "fmt"

func method(args ...interface{}){
	fmt.Printf("args is %T,values is %v" ,args,args)
}

func main(){
//空接口是万能类型，保存任意类型的值
var i interface{}=1
fmt.Println("i=",i)

i=[4]int{1,2,3}
	fmt.Println("i=",i)
method(i,123)

}
