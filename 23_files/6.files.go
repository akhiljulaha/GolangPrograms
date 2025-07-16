// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Reading data from the first file and then insert into the second file 
// 		// ✅ Step 1: Open the source file (example.txt)
// sourceFile, err := os.Open("example.txt")    // opening the source file 
// if err != nil{
// 	panic(err)
// }

// defer sourceFile.Close()           // defer and close 

// 	// ✅ Step 2: Create the destination file (example2.txt)
// destFile, err := os.Create("example2.txt")    
// if err != nil{
// 	panic(err)
// }
// defer destFile.Close()

// 	// ✅ Step 3: Create buffered reader and writer for efficient I/O
// reader := bufio.NewReader(sourceFile)    // add it
// writer := bufio.NewWriter(destFile)     // add it 

// 	// ✅ Step 4: Copy byte-by-byte until EOF (end of file)
// for{                             
// 	b, err := reader.ReadByte()
// 	if err != nil{
// 		if err.Error() != "EOF"{
// 			panic(err)
// 		}
// 		break
// 	}

// 	e := writer.WriteByte(b)        // write 
// 	if err != nil{
// 		panic(e)
// 	}
// }
// 	// ✅ Step 5: need to flash the data 

// writer.Flush()
// fmt.Println("written to new file successfully")
	
// }

