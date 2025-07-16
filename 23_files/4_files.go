package main

import (
	"fmt"
	"os"
)

func main() {


	// ✅ Reading files/folders in a directory

	// "."   = current folder
	// ".."  = parent folder (one level up)
	dir, err  := os.Open("../")   // Open the parent folder
	if err != nil{
		panic(err)
	}
	defer dir.Close()

	// fileInfo, err := dir.ReadDir(6)    //  will only return 6 entries max 
	fileInfo, err := dir.ReadDir(-1)       //  Use -1 to get **all files and folders**


	// Loop through the entries and print their names
	for _,fi := range fileInfo{
		fmt.Println(fi.Name())
	}
	
}


	//2- read file -- 3 WAY (if the file is heavy)

	// data, err := os.ReadFile("example.txt")
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(string(data)