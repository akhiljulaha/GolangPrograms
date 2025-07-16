// package main
// import (
// 	"fmt"
// )
// /*   5
// Unbuffered channels:
// - Both sending and receiving are blocking operations.
// - So if one side (send or receive) isn’t ready, the program will deadlock.

// Buffered channels:
// - Allow sending values without blocking until the buffer is full.
// - Useful when receiver isn’t immediately ready.

// */

// func main() {

// 	// messageChan := make(chan string) 
// 	// messageChan <- "ping" 
// 	// msg := <-messageChan 
// 	// fmt.Println(msg)

// 	// 	💡 Solutions:
// 	// 1. Use a goroutine to receive/send in parallel (for unbuffered channel).  (using "Go" keyword)
// 	// 2. Use a buffered channel so send doesn't block immediately.(below one)

// 	emailChan := make(chan string, 100)  // buffer capacity
// 	emailChan <- "1@example.com"
// 	emailChan <- "2@example.com"
// 	fmt.Println(<-emailChan)
// 	fmt.Println(<-emailChan)

// 	// here you can see there is  no deadlock condition here because using the Buffered channel approach


	
// }

 