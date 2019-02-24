package main

import (
	"regexp"
	"fmt"
)

func main(){
	buf:="291923.10293 129900 0129929 29123% 10101010  00000 5.6 7.8 55.7"
	reg1:=regexp.MustCompile(`\d\.\d`)
	if reg1==nil{
		fmt.Println("error")
	}
	result:=reg1.FindAllString(buf,-1)
	fmt.Println(result)

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1:=data[8:]  //[8 9]
	s2:=data[:5]  //[0 1 2 3 4]
	s2[4]=100
	fmt.Println(s2)
	fmt.Println(s1)
	copy(s2, s1)
	fmt.Println(data)

}
