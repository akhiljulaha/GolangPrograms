// package main
// import (
// 	"fmt"
// )
// // Now in Step 4:
// // 🔁 Goroutine is the sender → sending result from sum() function
// // 📥 Main is the receiver → reading the result in main using <- result

// func sum(result chan int, num1 int, num2 int){
// 	numResult := num1 + num2
// 	result <- numResult   // sending inside the channel
// }
// func main() {


// 	result := make(chan int)

// 	go sum(result, 4,5)
// 	res := <- result    // receiving the data from the channel
// 	fmt.Println(res)





// }

 