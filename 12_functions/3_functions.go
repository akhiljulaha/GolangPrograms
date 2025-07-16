// // ✅ 3. Passing a Function as a Parameter
// package main

// import "fmt"

// // ➤ 3. Function that accepts another function as a parameter
// func processIt(fn func(a int) int) {
// 	result := fn(5)
// 	fmt.Println("Processed result:", result)
// }  

// func main() {
// 	fn := func(a int) int {    // Anonymous Function 
// 		return a * 2   
// 	}		
// 	processIt(fn)
// }

// // 1) Creating an anonymous function inside main() and assigning it to the variable fn.
// // 2) Calling processIt(fn) → passing the function as an argument.
// // 3) Inside processIt, result := fn(5) → this calls the function passed from main() with the value 5 and returns the result (5 * 2 = 10).




// // ✅ In Golang, we can create anonymous functions inside the main() function 
// // and pass them to other functions like processIt().
// // ❌ In Java, we cannot define named functions inside the main() method, 
// // and we also cannot call a function defined inside main() from another class-level method.