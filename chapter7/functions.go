package main 

import "fmt"


func main(){
	fmt.Println(average([]float64{98,93,77,82,83}))
	fmt.Println(multiValue(10,5))
	fmt.Println(variadic(1,2,3))
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	panic("PANIC")

}

// simple function returing average of input array
func average(elements []float64) float64 {
	sum := 0.0
	for _,value:=range elements{
		sum+=value
	}
	return sum / float64(len(elements))
}

// return value using return variable name
// func average(elements []float64) (avg float64) {
// 	sum := 0.0
// 	for _,value:=range elements{
// 		sum+=value
// 	}
// 	avg = sum / float64(len(elements))
// 	return 
// }

// function returing multiple values
func multiValue(a int,b int) (int,int){
	return b,a
}

// variable length argument functions
func variadic(args ...int) int{
	total := 0
	for _,value := range args{
		total+=value
	}
	return total
}

func makeEvenGenerator() func() int {
     i := 0
     return func() (ret int) {
          ret = i
		  i += 2
		return 
	}
}


