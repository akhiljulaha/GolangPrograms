// package main
// import (
// 	"fmt"
// 	"time"
// )
// // order struct – like a class in Java, but with no methods
// type order struct{
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time   // nanosecound precision
// }  
// func main(){
// 	// Creating an instance (object) of the struct – similar to creating an object from a class in Java
// 	myOrder := order{              //1
// 		id: "1",
// 		amount: 50.00,
// 		status: "received",
// 	}
// 	myOrder.createdAt = time.Now()    // later on also we can add
// 	fmt.Println(myOrder.status)   // want to fetch the specific value
// 	fmt.Println("Order struct", myOrder)                   

// // Creating another order instance with all fields initialized
// 	myOrder2 := order{             //2
// 		id: "2",
// 		amount: 100,
// 		status: "delivered",
// 		createdAt: time.Now(),
// 	}
// myOrder.status = "paid"       // updating 
// 	fmt.Println("Order struct", myOrder2)                   
// 	fmt.Println("Order struct", myOrder)                   


// }
// // Note:
// // - Structs in Go are similar to classes in Java but without inheritance or methods.
// // - You can create multiple instances (like objects).
// // - Structs support encapsulation and can work with methods via receiver functions, which supports basic OOP in Go.  