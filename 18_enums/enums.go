package main

import "fmt"

// Enumerated types
// Creating Custom Type

// type OrderStatus int      //1-a   // Define a new custom type called OrderStatus based on int
// // Decla			re a list of constant values of type OrderStatus using iota
// const(                             // 1-b
// 	Received OrderStatus = iota   // iota=0 it's untyped integer and i will increment
// 	Confirmed
// 	Prepared
// 	Delivered
// )
type OrderStatus string               // 2-a
const(                             // 2-b
	Received OrderStatus = "received"   
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)


// Define a function that accepts an OrderStatus type as parameter
func changeOrderStatus(status OrderStatus){                     //1-c
	fmt.Println("Chnaging order status to", status)
} 

func main(){
	changeOrderStatus(Confirmed)                        // enums used when mutiple values need to pass


}

// why Enums
// -> facing typo issue 
// ->