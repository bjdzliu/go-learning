package main

import (
	"fmt"
)

type Address struct{
	addr string
}

type Vcard struct{
	addr *Address
	name string
}

func main(){
	user1:=Vcard{&Address{"beijing"},"xiaoming"}
	fmt.Println("user1 info :",user1)
	fmt.Println("user1 info :",*(user1.addr))
	fmt.Println("user1.addr info :",user1.addr)

	addr1:=Address{"haidian"}
	fmt.Printf("content is %s\n",addr1.addr)

	user2:=Vcard{addr:&addr1}
	fmt.Println("user2",user2)

}