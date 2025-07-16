// package main

// import (
	
// 	"os"
// )

// func main() {


// 	// ✅ Create (or overwrite) a file named "example2.txt"
// 	// If the file already exists, it will be truncated (cleared).
	
// 	f, err  := os.Create("example2.txt")   // Open the parent folder
// 	if err != nil{
// 		panic(err)
// 	}
// 	defer f.Close()

	
// 	// f.WriteString("Hi Go")         	// ✅ Option 1: Write a simple string
// 	// f.WriteString("nice Go")   // will append/add the file, will not replace the privious one 


// 	bytes := []byte("hello golang")   	// ✅ Option 2: Write using byte slice
// 	f.Write(bytes)
// }

