package main

import (
	"fmt"
)

func main() {
	i1 := int(1)
	i2 := int32(1)
	i3 := int64(1)

	p1 := &i1
	p2 := &i2
	p3 := &i3

	p1 = (*int)(p2)
	//	p1 = (*int)(p3)

	//	p2 = (*int32)(p1)
	//	p2 = (*int32)(p3)

	//	p3 = (*int64)(p1)
	//	p3 = (*int64)(p2)

	//	p1 = p1 + 8

	fmt.Printf("p1=%p,p2=%p,p3=%p\n", p1, p2, p3)
}
