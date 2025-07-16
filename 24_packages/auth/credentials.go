package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username,password)
}
// Function names starting with a capital letter are exported and can be accessed from other packages.
// If the function name starts with a lowercase letter, it is unexported and accessible only within the same package.	