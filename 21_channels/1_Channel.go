// package main

// import (
// 	"fmt"
// 	"time"
// )

// func processNum(numChan chan int){
// 	fmt.Println("processing number", <-numChan)   // reciving 
// }

// func main() {

// // 2 we are passing a value from one goroutine (main) to another goroutine (processNum) using a channel.
// 	numChan := make(chan int)    

// 	go processNum(numChan)      //IMP -> why it gets STUCK without go -->refer the Notes

// 	numChan <- 5   

// 	time.Sleep(time.Second*2)
// 	/*
// // ✅ What’s happening? (IMP to undersatend)
// // -The line go processNum(numChan) starts a new goroutine (background thread).
// // -That goroutine waits to receive a value from the channel.
// // -Then, in the main function, we send a value with numChan <- 5.
// // -Since the receiver is already waiting in another goroutine, the send doesn’t get blocked, and everything works fine.
// // 	*/


	
	
// 	//  1 step : to check the deadlock condition
// 	// messageChan := make(chan string) // created the channel

// 	// messageChan <- "ping" // sending the data inside the channel(messageChan)

// 	// msg := <-messageChan // reciving the data from the channel
// 	// fmt.Println(msg)

// }

