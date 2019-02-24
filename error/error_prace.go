package main

import (
	"fmt"
	"errors"
)

func main(){
	err1:=fmt.Errorf("%s","this a normal")
	fmt.Println(err1)
	err2:=errors.New("this is a normal err2")
	fmt.Println(err2.Error())

result,err:=myDiv(10,0)
if err!=nil{
	fmt.Println("err=",err)
}else{
	fmt.Println(result)
}

}

//err是接口类型
func myDiv(a,b int )(result int,err error){
	err=nil
	if b==0{
		err=errors.New("分母不能是0")
	}else {
		result=a/b
	}
	return result,err

}
