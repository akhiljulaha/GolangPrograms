package main
import "fmt"

func changeNum(num int) {                   //1  call by value 
	num = 5
	fmt.Println("In ChangeNum", num)
} 

// func changeNum(num *int) {          //2  call by referance  ,* is defne address of the variable
// 	*num = 5      
// 	fmt.Println("In ChangeNum", *num)
// }
func main() {
num := 1
changeNum(num)           // 1a
// changeNum(&num)                                   //2a
// fmt.Println("After changeNum in main", num)      

// fmt.Println("Memory address ", &num)  // &num helps  to find memory address, here the num variable is stored

}