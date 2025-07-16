//✅ 4. Returning a Function from Another Function
package main

import "fmt"

// ➤ processIt returns an anonymous function that accepts an int and returns its square.
// This returned function can be stored and called later with a value.
func processIt() func(a int) int {
	return func(a int) int {
		return a * a           //it returns a function that can later accept a value.
	} 
} 
func main() {
	fn := processIt()//1) it returns => func(a int) int { return a * a }, So now fn holds a function that expects an integer.
	result := fn(2)  // 2 we are executing the function stored in fn with the argument 2.
	fmt.Println("Returned function result:", result)
}

// Note:
// processIt() itself doesn't receive any value.
// It simply returns a function.
// That returned function is executed later with an argument (e.g., 2),
// and that's when the calculation a * a is performed → 2 * 2 = 4
