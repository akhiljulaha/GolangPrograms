package main
import (
	"fmt"
	"time"
)
// passing the struct inside other struct
type customer struct{
	name string
	phone string
}
 
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecound precision
	customer
}
func main() {
// how to set the customer details and need to pass in the order struck
newCustomer := customer{
	name: "john",
	phone: "1232345666",
}

newOrder := order{
		id : "1",
		amount: 30,
		status: "received",
		customer: newCustomer,      //IMP
}
		fmt.Println(newOrder)
		// newOrder.customer.name = "robin"
		// newOrder.amount = 88
		// fmt.Println(newOrder)

 

}