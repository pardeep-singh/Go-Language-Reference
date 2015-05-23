package math

func Average(elements[] float64)float64{
	total := float64(0)
	for _;x:=range elements{
		total+=x
	}
	return total/float64(len(elements))
}
