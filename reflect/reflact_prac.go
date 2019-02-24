package main

import (
	"reflect"
	"fmt"
)

type TagType struct{
	field bool "an"
	field2 string "tow"
	field3 int "sho"
}

func main(){

	tt:=TagType{true,"bbb",1}
	for i:=0;i<3;i++{
refTag(tt,i)
	}
}

func refTag(tt TagType,ix int){
	ttType:=reflect.TypeOf(tt)
	fmt.Println("self define struct :",ttType)
	ixField:=ttType.Field(ix)
	fmt.Printf("%v\n",ttType.Key)
	fmt.Printf("%v\n",ixField.Tag)
}