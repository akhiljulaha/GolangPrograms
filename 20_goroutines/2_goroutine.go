// package main
// import (
// 	"fmt"
// 	"sync"
// )
// func task(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() // 🔽 Step 3: Tell Go this goroutine is finished
// 	fmt.Println("Doing task", id)
// } 
// func main() {
// 	var wg sync.WaitGroup // 🔁 Step 1: Create a WaitGroup

// 	for i := 0; i <= 10; i++ {
// 		wg.Add(1)             // ➕ Step 2: Tell Go "I'm launching 1 new goroutine"
// 		go task(i, &wg)       // 🚀 Launch goroutine and pass WaitGroup pointer
// 	}
// 	wg.Wait() // ⏳ Step 4: Main waits here until all `Done()`s are called
// }
