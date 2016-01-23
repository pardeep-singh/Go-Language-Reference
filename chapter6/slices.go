package main
import "fmt"

func main() {
	baseArray := [5]int{1,2,3,4,5}
//	var slice[] int

//	slice =  make([]int,3)
//	fmt.Println(len(slice))

	slice := baseArray[0:2]
	slice[0]=9

	slice1 := baseArray[2:3]

	// slice = append(slice1,slice)
	slice = append(slice1,1,2)

	copiedSlice := make([]int,2)
	copy(copiedSlice,slice)
	copiedSlice[0]=99

	fmt.Println(slice[0]," ",baseArray[0])

	fmt.Println(slice," ",slice1," ",baseArray," ",copiedSlice)
	fmt.Println(len(slice)," ",len(slice1)," ",len(baseArray)," ",len(copiedSlice))

	slice = append(slice,13,14)
	fmt.Println(slice," ",slice1," ",baseArray," ",copiedSlice)
	fmt.Println(len(slice)," ",len(slice1)," ",len(baseArray)," ",len(copiedSlice))

}