package main

import "fmt"

func MinAndMax(num1 int, num2 int) (min int, max int) {

	if num1 > num2 {

		min = num2

		max = num1

	} else {

		max = num2

		min = num1

	}
	return

}
func main() {
	min,max := MinAndMax(33, 22)
	fmt.Printf("min = %d, max = %d\n", min,max)
}
