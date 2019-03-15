package main

import (
	"strconv"
	"fmt"
)


func main() {
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`) // there is a tab character inside the string literal
	fmt.Println(s)

}
