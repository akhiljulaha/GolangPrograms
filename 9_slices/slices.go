package main

import (
	"fmt"
	"slices"
)

// slice --> dynamic array
// -> Most used construct in golang
// -> + useful methods

func main(){
 
// ➤ 1. Nil Slice (Uninitialized)
var nums []int  // (Declared but not initialized)similar to array,not mentioning the size
	fmt.Println(nums== nil)  //true → It's nil (no memory allocated)
	fmt.Println(len(nums))    //0 → No elements

// ➤ 2. Make Slice (Initialized with make)
	// var nums = make([]int,2,5)    // Length = 2 → [0, 0], Capacity = 5 → space for 3 more elements
	// fmt.Println(nums== nil)   //false
	// fmt.Println(nums)
// capacity -> maximum numbers of elements can fit(2)
	// fmt.Println((cap(nums)))  

    // ➤ 3. Append Elements
// var nums = make([]int,2,5)	
// nums = append(nums,100)
// nums = append(nums,200)
// nums = append(nums,300)
// nums = append(nums,400)
// capacity will double once reach the limit
// fmt.Println(nums)
// fmt.Println((cap(nums)))      // Likely 10 (capacity doubles)
// fmt.Println(len(nums))   


// ➤ 4. Another Way to Create Slice (Literal)
// here capacity is getting updated as per your adding values
// nums := []int{}
// nums = append(nums,100)
// nums = append(nums,200)

// fmt.Println(nums)
// fmt.Println((cap(nums))) 
// fmt.Println(len(nums)) 


// ➤ 5. Index Assignment	
// var nums = make([]int,2,5)	
// nums[0] =3
// nums[2] =4    // ❌ Invalid: index out of current length

// fmt.Println(nums)
// fmt.Println((cap(nums))) 
// fmt.Println(len(nums))

// 6. Copying a Slice (using built-in copy function)
// var nums = make([]int,0,5)   // L->0 , C-> 5
// fmt.Println(nums)          //[]
// nums = append(nums,2)      // L-> 1

// var nums2 = make([]int, len(nums))// Create a new slice with the same length (L->1 but it 0)
// copy(nums2,nums)     	               // Copy elements from nums → nums2
// fmt.Println(nums, nums2)


//7. Slice Operator (Slicing a slice)

// var nums3 = []int{1,2,3}
// fmt.Println(nums3[0:2]) // [1 2] → end index is excluded
// fmt.Println(nums3[:2])  // [1 2] → from beginning to index 2 (excluded)
// fmt.Println(nums3[1:])  // [2 3] → from index 1 to end


// 8. slices package (Go 1.18+)
var nums1 = []int{1,2,3}
var nums2 = []int{1,2}
fmt.Println(slices.Equal(nums1,nums2))


// 9. 2-D Slices (Slice of Slices)
// var matrix = [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 	}

// 	fmt.Println(matrix) // [[1 2 3] [4 5 6]]



 




}


// 📝 Recap:
// ➤ Length (len)    = number of elements currently visible in the slice
// ➤ Capacity (cap)  = total space available in the backing array