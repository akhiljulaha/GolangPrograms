package main
import "fmt"
//interface
type Voice interface{   // 1. Define interface with Say() method
	Say() string //when using interface as type so make sure inside only declaration hapeen not implementaion
}
type Cat struct{}             //  2. Cat struct 

func (c Cat) Say() string{      // 3. Cat implements Voice by defining Say()
	if 1 > 0{
		return "cat	 something"
	}
	return "meow"
}
type Dog struct{}                // Dog struct implementing Voice interface

func (d Dog) Say() string {
	 return "dog woof"
}
func main() {
	c := Cat{}   // 4. Create Cat instance
	d := Dog{}
	voice := []Voice{c,d} //5 IMP->created a slice of type Voice,This allows storing any type that implements Say()
	for _,v := range voice{     // 6. Loop through interface slice and call Say()
		fmt.Println("voice of ", v.Say())
	}
}
// Notes:
// - Interfaces can be used in two ways:
//   1. As a data type (e.g., slice of Voice)
//   2. To define behavior (methods) for cleaner and modular code
//IMP It is mandatory that both structs must implement the same method(s) defined in the interface — with the same method signature. 