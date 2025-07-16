// package main
// // like in java we are using constructor to intialize the class values
// // the same will do in the struct. basically in the golang don't have contructor approach but hack we can use

// import (
// 	"fmt"
// 	// "time"
// )
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	// createdAt time.Time // nanosecound precision
// }
// func newOrder(id string, amount float32, status string) *order{     //pointer of the order struct
// myOrder := order{              
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder  // we can't return directly, returning the pointer (&)
// }
// func main() {
// 	// myOrder := order{ // whatever we are doing inside the main the same will do inside    newOrder(constructor)                       //1-Direct way
// 	// 	id: "1",
// 	// 	amount: 50.00,
// 	// 	status: "received",
// 	// }
// 	// fmt.Println(myOrder)  

// 	myOrder := newOrder("1", 30.50,"recived")
// 	fmt.Println(myOrder)  
// 	fmt.Println(myOrder.amount)  

// } 
	


// // In direct struct assignment, if you skip fields, Go just uses the zero value (e.g., "", 0, false) — no compile-time error.
// // In a constructor function (newOrder(...)), you control what values are required — so the user must pass all required arguments, ensuring safety and completeness.
// // ✅ Constructor approach is better when you want to enforce mandatory values during struct creation.


