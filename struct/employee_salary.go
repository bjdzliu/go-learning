package main

type employee struct{
	salary int
}

var count float32

func (this *employee) giveRaise ( percent float32) {
count=float32(this.salary)*percent
}
func main(){

}