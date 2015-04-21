package main

import "fmt"

func main(){
	elementsInfo := map[string]map[string]string{
		"H":map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He":map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li":map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
		"Be":map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B":map[string]string{
			"name":"Boron",
			"state":"solid",
		},
	}

	for outerKey := range elementsInfo{
		fmt.Print("Outer Key:",outerKey)
		for innerKey := range elementsInfo[outerKey]{
			fmt.Print(" Inner Key:",innerKey," Value:",elementsInfo[outerKey][innerKey])
		}
		fmt.Println("\n")
	}
}