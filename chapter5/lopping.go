package main

import "fmt"

func main(){
	i:=0
	for i<=10 {
		fmt.Print(i," ")
		i+=1
	}
	fmt.Println()
	for j:=1;j<=10;j++{
		if (j%2==0) {
			fmt.Print(j," Even ")
		} else { 
			fmt.Print(j," Odd ")
		}
	}
	fmt.Println()

	var j int
	fmt.Scanf("%d",&j)
	switch j{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("default")
	}
}