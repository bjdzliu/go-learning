package main

import (
	"io"
	"bytes"
	"fmt"
)

func main(){

	var w io.Writer=new(bytes.Buffer)
	fmt.Printf("%T",w)
}
