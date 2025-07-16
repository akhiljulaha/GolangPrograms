package main

import "fmt"

func main() {
	var nums [4]int      // here just alocating(int ---> imp)
// array length
	fmt.Println(len(nums))   

// add value in the array
	// nums[0] = 1
	// fmt.Println(nums[0])

// print the complete array
	// fmt.Println(nums)


// booleaan case
	// var vals[4]bool
	// vals[2] = true
	// fmt.Println(vals)  // in case of boolen then bydefault will be false(int -> 0)

// String case

// var names[4]string
// names[1] = "akhil"
// fmt.Println(names) 

// NOTE:   int -> 0,   string -> " ", bool -> false


// dectare it in the single line 
// nums := [3]int{1,2,3}
// fmt.Println(nums)


// 2-D array
// nums := [2][2]int{{3,4},{5,6}}
// fmt.Println(nums)

// Benefits
// - Fixed size, that is predictable
// - Memory optimization
// - Constant time access(becusse we know the index)

//----> mostly will use the slices instead of the array
// array -> static memory     slices -> dynamic memory

} 