// package main

// import (
// 	"fmt"
// )

// // type stack struct{
// // 	elements []int   // this will works only for the integer, below is generic approach   //1
// // }
// type stack[T any] struct{
// 	elements []T
// }
 

// func main() {
// 	myStack := stack[int]{
// 		elements: []int{1,2,3,5},
// 	}
// 	fmt.Println(myStack)

// }

 