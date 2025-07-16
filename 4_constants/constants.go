package main

import "fmt"

const age = 30    //If declaring the variable outside a function then it's fine if not using it 
// name := ""  // can't use the shorthand outside the function
var name string = "test"

func main() {

	const name string = "golang"

	// name = "javascript"

	fmt.Println(name)

	// Multiple constants using at the same time
	const( // only 1 also we can use, there is no restriction need to use all inside the const
		port = 5000
		host = "loacalhost"

	)

	fmt.Println(port,host)

}

// Constant value can't be changed
// 

// It is used to declare and initialize the variables only inside the functions.
 