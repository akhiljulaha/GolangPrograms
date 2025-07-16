// package main
// import "os"

// func main() {

// 	// 2 - Read file (WAY 1): Open the file "example.txt" for reading.
// 	f, err := os.Open("example.txt")
// 	if err != nil{
// 		panic(err)
// 	}

// 	defer f.Close()   	// Always close the file after you're done. 'defer' ensures it runs at the end of main().
// 	// storing the file data 
// 	buf := make([]byte,13)   	// Create a byte slice (buffer) of size 13 to store the data read from the file.

	
// 	// Read up to 13 bytes from the file and store it into the buffer.
// 	// 'd' contains the actual number of bytes successfully read.
// 	d, err := f.Read(buf)   // reading the data 
// 	if err != nil{
// 		panic(err)
// 	}
// 	// Loop through each byte and print it as a character (string).
// 	// Note: This prints each character one by one.

// 	for i:=0;i<len(buf);i++{
// 	// println("data",d, buf)   // make sure converting into string
// 	println("data",d, string(buf[i]))
// }

// }
// 	// ✅ Bonus Tip: Instead of printing each character one-by-one,
// 	// you can also print the whole buffer using:
// 	// println("Full data:", string(buf))