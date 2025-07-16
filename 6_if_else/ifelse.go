package main

import "fmt"

func main() {

//-> if-else
	// age := 13
	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// }else{
	// 	fmt.Println("person is not an adult")
	// }

//-> else-if
	// age := 1
	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// }else if age >=12{
	// 	fmt.Println("person is teenager")
	// }else{
	// 	fmt.Println("person is a kid")
	// }

// logical way 
	// var role = "admin"
	// var hasPermissions = false
	// if role == "admin" && hasPermissions{
	// 	fmt.Println("yes")
	// }else{
	// 	fmt.Println("No")
	// }

// we can declare a variable inside the if construct	
	if age:=20; age>=18{
		fmt.Println("Person is an adult", age)
	}else if age >= 12{
		fmt.Println("Person is teenager", age)
	}

// Note: golang does not have ternary, we have to use the normal normal if else concept


} 