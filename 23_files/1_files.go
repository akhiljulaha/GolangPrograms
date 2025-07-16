// package main

// import (
// 	"fmt"
// 	"os"
// )
// func main() {
// 	// Open the file named "example.txt" in read-only mode.
// 	// 'f' is the file object, 'err' captures any error that might occur while opening.
// 	f,err := os.Open("example.txt")
// 	if err != nil{
// 		 // If there's an error (e.g., file not found), panic will stop the program and print the error.
// 		panic(err)  // For now, use panic to understand error handling.
// 	}

// // Get detailed information about the file using the Stat() method.
// // It returns fileInfo, which holds metadata like name, size, permissions, etc.
// 	fileInfo, err := f.Stat()
// 	if err != nil{
// 		//log the error
// 		panic(err)   
// 	}
// 	// Print the file name (just the name, not full path)
// 	fmt.Println("File name:", fileInfo.Name())

// 	// Check whether it’s a file or a directory (folder)
// 	fmt.Println("Is it a folder?:", fileInfo.IsDir())

// 	// Print the size of the file in bytes
// 	fmt.Println("File size (in bytes):", fileInfo.Size())

// 	// Print file permissions (like read/write access)
// 	fmt.Println("File permissions:", fileInfo.Mode())

// 	// Print the last modified time of the file
// 	fmt.Println("Last modified:", fileInfo.ModTime())







// }