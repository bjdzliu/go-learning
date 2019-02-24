package external_p2
import (
	"helloworld/external_package"
	"fmt"
)
func Output2(){
	fmt.Println("in output2")
	external_package.Setvalue()
	external_package.Out()
}

