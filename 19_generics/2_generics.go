// package main

// import (
// 	"fmt"
// 	// "runtime/trace"
// )
// func PrintSlice[T any](items []T) { //will allow any data type,for restriction use the  below one  //1
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// // func PrintSlice[T int | string| bool](items []T) {  
// // 	for _, item := range items {
// // 		fmt.Println(item)
// // 	}
// // }

 
// func main() {
// 	// nums := []int{1,2,3}   //1 
// 	names := []string{"golang", "typescript"}
// 	PrintSlice(names)

// 	// vals := []bool {true,false}    // this will not alow at the compiler time
// 	// PrintSlice(vals)

// }

 