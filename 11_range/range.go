package main

import "fmt"

// iterating over data structure

func main(){
	// nums := []int {6,7,8}   // slice

	//1- iterating 
	// for i:=0; i< len(nums); i++{
	// 	fmt.Println(nums[i])
	// }

	// ➤ 2. Range with index and value

	// nums := []int {6,7,8} 
	// for i, num := range nums{
	// fmt.Println(num , i)
	// }

	//3- sum 
	// nums := []int {6,7,8} 
	// sum := 0
	// for _, num := range nums{
	// 	sum = sum + num
	// }
	// fmt.Println(sum)



// ➤ 4. Range over map (key + value)
	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// ➤ 5. Range over map (key only)
	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for k := range m {
	// 	fmt.Println(k)
	// }


// ➤ 6. Range over string (unicode rune values)

	for i, c := range "golang"{
		fmt.Println(i, c)
	}

// ➤ 7. Print string characters properly (convert rune → string)

	for i, c := range "golang"{
		fmt.Println(i, string(c))
	}


  






}