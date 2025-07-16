// package main
// import (
// 	"fmt"
// )
// /* Step 5  (Goroutine synchronizer)
// Earlier, we used WaitGroup to wait until all goroutines completed before the main function ended.
// Now, we use a Buffered Channel to collect signals from each goroutine — and wait by receiving those values in main.
// */

// func task(done chan bool){
// 	defer func ()  {done <- true}()//you can skip defer — just ensures that done <- true always runs, even if the function returns early or errors occur.
// 	fmt.Println("Processing....")
// 	// done <- true   // sending inside the done channel (might give error)
// }

// func main() {

// 	done := make(chan bool)	

// 	go task(done)
// 	<- done // This line blocks because it’s waiting to receive a value from the channel.

// //---> If the channel is unbuffered, <-done will block until someone sends a value.
// //---> When task() sends done <- true, it unblocks this line.

// // In an unbuffered channel, both sending and receiving are blocking operations.
// // To avoid this we can use a buffered channel.

	
// }

