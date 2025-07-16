// package main

// import (
// 	"fmt"
// 	"sync"
// )
 
// func worker(i int, wg*sync.WaitGroup){
// 	defer wg.Done()     // Step 3: Mark this worker as done in the WaitGroup
// 	fmt.Printf("worker %d started\n",i)
// 	// some task is happening
// 	fmt.Printf("worker %d end\n",i)

// }
// func main() {
// 	fmt.Println("Explore goroutine started")
// 	var wg sync.WaitGroup // Step 1: it's kind of checklist, Create a WaitGroup to wait for all workers
// 	// start  workers goroutines
// 	for i:=1; i<=3;i++{
// 		wg.Add(1)      // Step 2: Tell the WaitGroup we’re adding a new worker
// 		go worker(i, &wg)
// 	}
// 	wg.Wait() // Step 4: Wait until all workers have called Done()
// 	fmt.Println("workers task completed")

// }