package main

import "fmt"

//...int means the function can accept any number of int arguments — even zero.
// Internally, nums becomes a slice of integers ([]int).
func sum(nums ...int)int{
	total :=0
	for _,num:=range nums{ //range gives index and value, _ is used to ignore the value like ignoring the index
		total = total + num
	}
	return total
}
func main() {
 
	result := sum(1,2,3,4,5)      // 1 way 
	fmt.Println(result)

	// nums := []int{3,4,5,6}              // silce    2 way
	// result := sum(nums...)
	// fmt.Println(result)


}