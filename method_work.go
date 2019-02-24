package main

import (
	"fmt"
)

type TZ int

func (tz *TZ) Increate(a int) {
	*tz += TZ(a)
}

func main() {
	var a TZ
	//100是int，强制类型转换
	a.Increate(100)
	fmt.Println(a)

}
