package main

import (
	"encoding/json"
	"fmt"
)

func main(){

	m:=make(map[string]interface{},3)
	m["company"]="liudz"
	m["age"]=100
	m["addr"]="beijing"

	result,err:=json.Marshal(m)
	if err!=nil{
		fmt.Println("error append")
	}
	fmt.Printf("result's value is %T\n",result)
	fmt.Println(string(result))

}
