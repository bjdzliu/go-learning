package external_package

/*
used by map_prac.go

 */

import "fmt"

var exter_package = 1
func Out(){
	fmt.Println("print external package")
	fmt.Println(exter_package)
}

func Setvalue() {
	exter_package= exter_package+100
}