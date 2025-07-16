package main

import (
	"fmt"
	"os"
)

func main() {


	// want to delete the file 


err := os.Remove("example2.txt")   // there is one uniqe thing in the go we returning the error 
if err != nil{
	panic(err)
}

fmt.Println("file deleted successfully")


	
}

