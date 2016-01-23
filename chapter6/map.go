package main 

import "fmt"

func main() {
	chemicalElements := make(map[string]string)
	chemicalElements["H"] = "Hydrogen"
	chemicalElements["HE"] = "Helium"
	chemicalElements["LI"] = "Lithium"
	chemicalElements["BE"] = "Beryllium"
	chemicalElements["B"] = "Boron"
	chemicalElements["C"] = "Carbon"

	for eleKey := range chemicalElements{
		fmt.Println("Key:",eleKey," Value:",chemicalElements[eleKey])
	}

	element,status := chemicalElements["H"]
	fmt.Println("Status:",status," Value",element)

	if element,status := chemicalElements["H"];status{
		fmt.Println("Status:",status," Value",element)
	}
}
