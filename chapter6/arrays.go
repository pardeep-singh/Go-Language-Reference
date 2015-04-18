package main 

import "fmt"

func main() {
	
	{
		var floatArray [5] float64
		floatArray[0]=1
		floatArray[1]=2
		floatArray[2]=3
		floatArray[3]=4
	
		var elementsSum float64 = 0

		for i:=0;i<5;i++{
			elementsSum += floatArray[i]
		}
		fmt.Println(elementsSum/float64(len(floatArray)))
	}
	
	{
		floatArray := [5]float64{1,2,3,4,5}
		var elementsSum float64 = 0
		for _,arrayElement := range floatArray{
			elementsSum += arrayElement
		}
		fmt.Println(elementsSum/float64(len(floatArray)))
	}
}