package main

import "fmt"
 
func PrintSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func PrintStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}
func main() {
	nums := []int{1,2,3}   //1 
		PrintSlice(nums)

	// names := []string{"golang", "typescript"}
	// PrintStringSlice(names)

}

// The problem is that the same function logic is being repeated for different data types (like int and string), which leads to code duplication and makes the code harder to maintain and scale. If more data types are needed, separate functions must be created for each, even though the logic is identical. This violates the DRY (Don't Repeat Yourself) principle.

// The solution is to use generics, which allow writing a single function that can work with multiple data types — improving code reusability and maintainability. 