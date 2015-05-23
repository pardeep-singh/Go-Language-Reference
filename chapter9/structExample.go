package main
import "fmt"
import "math"

type Circle struct{
	x, y, r float64
}

//Normal function
func area(c Circle) float64{
	return math.Pi*c.r*c.r
}

//Receiver function

func (c *Circle ) area()float64{
	return math.Pi*c.r*c.r
}

func main(){
	c := Circle{0,0,5}

	fmt.Println(area(c))

	fmt.Println(c.area())
}