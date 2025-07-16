package main

import (
	"fmt"
	// "os/user"

	"github.com/akhil/podcast/auth"
	"github.com/akhil/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("test", "secret")
	session := auth.GetSession()
	fmt.Println("session", session)

user := user.User{
	Email: "user@gmail.com",
	Name: "john joe",
}
// fmt.Println(user.Email, user.Name)
color.Red(user.Email)

}

