package main

import "fmt"

func max(num1,num2 int)int{
	var max int
	if num1<num2{
		max=num2
	}else if num1==num2{
		max=0

	}else if num1>num2{
		max=num1
	}
	return max

}

func main(){
	result:=max(100,999)
fmt.Println("max value is result",result)

}