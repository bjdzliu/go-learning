package main
import "fmt"

func main(){

	//类型相近的类型转换
	var a float32 = 1.1
	b := int(a)
	fmt.Println(b)
	c:=

	//string和int的转换
	var a6 int = 65  //大A的ASCII码是65
	b6 := string(a6) //string转换成文本，65表示A，所以b6就是A
	fmt.Println(b6)

}
