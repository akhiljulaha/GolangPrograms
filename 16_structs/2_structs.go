package main
import (
	"fmt"
	"time"
)
// attaching methods with structs like attching menthods with the class in java
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecound precision
}
//have to connect the function with struct with the helps of reciver type (o order)
func (o *order) changeStatus(status string) {
	if status != "confirmed" && status != "shipped" {
		fmt.Println("Invalid status")
		return
	}
	o.status = status
}

func (o order)getAmount()float32{      // if not changing the value then remove the star 
	return o.amount  
}
func main() {
	myOrder := order{              
		id: "1", 
		amount: 50.00,
		status: "confirmed",	
	}

	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder)  
	fmt.Println(myOrder.getAmount())                   
    myOrder.status = "passed"
	fmt.Println(myOrder.status)

 

} 
// ✅ Struct Field Access vs Method
//ans:  You can read or update struct fields directly:

// myOrder.amount = 700
// fmt.Println(myOrder.amount)
// But use methods when:
// ->You need to add validation
// ->Want to reuse logic


// //----> IMPORTANT
// // ✅ Direct field access (myOrder.status = "confirmed") is simple but doesn’t allow validation or logic control.
// // ✅ Using methods (like changeStatus()) is preferred when you want to enforce rules, reuse logic, or keep the struct safe from invalid updates.
// // ✅ Use *struct (pointer receiver) if the method needs to modify the original struct’s fields.



