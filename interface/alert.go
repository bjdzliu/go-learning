package main
//判断接口类型中所存的值类型
import "fmt"

type Student  struct{
	name string
	id int
}

func main(){
i:=make([]interface{},3)  //i是切片类型
i[1]=1
i[0]="ac"
i[2]=Student{name:"name",id:1919}

for _,valuei :=range i{
if value,ok:=valuei.(int);ok==true{
	fmt.Printf(" %v is int\n",value)
	}else if value,ok:=valuei.(string);ok==true{
	fmt.Printf(" %v is string\n",value)
}else if  value,ok:=valuei.(Student);ok==true{
	fmt.Printf(" %v is Student\n",value)
}
}

for index,data:=range i{
	switch value:=data.(type){
	case int:
		fmt.Printf(" %v is int,index is %d\n",value,index)
	case string:
		fmt.Printf(" %v is string,index is %d\n",value,index)
	case Student:
		fmt.Printf(" %v is string,Student is %d\n",value,index)
	}
}


}