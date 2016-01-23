package main
import "fmt"
func updateValue(value *int){
	*value += 1
}

func main(){
	
	//Call By Reference Example
	initialValue := 10
	fmt.Println("Initial Value:",initialValue)
	updateValue(&initialValue)
	fmt.Println("Updated Value:",initialValue)
	
	// Dynamically Allocated Pointer
	dynPointer := new(int)
	*dynPointer = 100
	fmt.Println("Initial Value:",*dynPointer)
	updateValue(dynPointer)
	fmt.Println("Updated Value:",*dynPointer)
}