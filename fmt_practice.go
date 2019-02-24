package main

import (
	"fmt"
)

func main() {
	var a int = 50
	var b bool
	c := 'a'
	type point struct {
		x, y int
	}
	p := point{1, 2}
	fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%v\n", p)


	fmt.Println(a, b, c)
	fmt.Printf("%T", c)
	fmt.Println()
	fmt.Printf("%t", b)
	fmt.Println()
	fmt.Printf("%v", b)
	fmt.Println()
	fmt.Printf("%b", 5)
	fmt.Println()
	fmt.Printf("%f\n", 2.1)
	fmt.Printf("%q\n", "")
	str := fmt.Sprintf("%q", "111")
	fmt.Println(str)

	var n int
	fmt.Scan("%d", &n)
	fmt.Print(n)

}
