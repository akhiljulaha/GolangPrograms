package main

import (
	"fmt"
	"maps"
)

// maps -> hash

// Key-value pairs | Keys are unique | Fast lookup
func main(){
	 
// ✅ 1. Create a map using make()
m := make(map[string]string)
// setting an element
m["name"] = "golang"
m["area"] = "backend"

// get an element
fmt.Println(m["name"], m["area"])

// try to use the key that is not exist
// fmt.Println((m["phone"]))   // if key does not exists in the map then it will return the zero (string -> empty string, Integer -> 0, bool -> false)


// ✅ 2. Integer map
// m := make(map[string]int)
// m["age"] = 30
// m["price"] = 300

// fmt.Println(m["age"])
// fmt.Println(m["phone"])   // getting 0 because value selected in the form of int
// fmt.Println(len(m))

// 🔸 Delete a key
// delete(m, "price")
// fmt.Println(m)

// 🔸 Clear the map (Go 1.21+)
// clear(m)
// fmt.Println(m)

// ✅ 3. Create map using literal (without make)
// m := map[string]int{"price":40, "phone": 3}
// fmt.Println(m)
// 🔹 Tip:
	// - Use `make()` when keys/values are dynamic (not known at compile-time)
	// - Use literals when values are static/known in advance


// ✅ 4. Check if key exists in map
// m := map[string]int{"price":40, "phone": 3}
// v will store the value → 40
// ok will be true → because "price" exists in the map.
// v,ok := m["price"]   // in the go, mupliple things we can return 
// fmt.Println(v)   
// if ok{
// 	fmt.Println("all Ok")
// }else{
// 	fmt.Println("not Ok")	
// }

// ✅ 5. Compare two maps (Go 1.21+)

// m1 := map[string]int{"price":40, "phone": 3}
// m2 := map[string]int{"price":400, "phone": 3}

// fmt.Println(m1 ==m2)   // ❌ Invalid: maps can't be compared directly
// fmt.Println(maps.Equal(m1,m2))


} 
