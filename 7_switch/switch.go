package main

import (
	"fmt"
	// "time"
)

func main() {

	// simple switch
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}
	
// Note: in the Golang, no need to write the break statement after every case, by default golang is handling


//->  Multiple condition switch
// switch time.Now().Weekday(){
// case time.Saturday, time.Sunday:
// 	fmt.Println("it's a weekend")
// default:
// 	fmt.Println("It's working")
// }

 
//-> Type switch

// whoAmI := func (i interface{})  {
// 	switch i.(type){
// 	case int: 
// 	fmt.Println("it's an integer")	
// case string:
// 	fmt.Println("it's a string")
// case bool:
// 	fmt.Println("it's a boolean")
// default:
// 	fmt.Println("others")
// }	
// }


// whoAmI("golang")
// whoAmI(5)
// whoAmI(true)
// whoAmI(4.444)

}
