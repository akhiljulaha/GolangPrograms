// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )
// // 🔁 Main goroutine was the sender → sending data into the channel
// // 📥 Goroutine (function) was the receiver → reading from the channel
// func processNum(numChan chan int){
// 	for num := range numChan{          // when using the range so no need to use the array 
// 	fmt.Println("processing number", num)   // reciving 
// 	time.Sleep(time.Second)
// 	}
// }
// func main() {
// // ✅ Step 3 – Passing Multiple Values Between Goroutines   
// // Sending
// 	numChan := make(chan int)    
// 	go processNum(numChan)  

// for{
// 	numChan <- rand.Intn(100)     // will create a random no.
// }

// 	// time.Sleep(time.Second*2)   // here no need of sleep why?

// // This is an example of continuous communication from one goroutine (sender) to another goroutine (receiver) using an unbuffered channel.




// 	/*
// ✅ What’s happening? (IMP to undersatend)
// -The line go processNum(numChan) starts a new goroutine (background thread).
// -That goroutine waits to receive a value from the channel.
// -Then, in the main function, we send a value with numChan <- 5.
// -Since the receiver is already waiting in another goroutine, the send doesn’t get blocked, and everything works fine.
// 	*/

// }

