package main

import(
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Does test contains es ? ->>>> ",strings.Contains("test","es"))

	fmt.Println("How many times t occur in test ? ->>>>",strings.Count("test","t"))

	fmt.Println("Does test has prefix te ? ->>>>",strings.HasPrefix("test","te"))

	fmt.Println("Where does s occur in tesst? ->>>>",strings.Index("tesst","s"))

	fmt.Println("Join a and b with sperator as - ->>>>",strings.Join([]string{"a","b"}, "-"))

	fmt.Println("Repeat a 5 times ->>>>",strings.Repeat("a", 5))

	fmt.Println("Replace a by b twice ->>>>",strings.Replace("aaaa", "a", "b", 2))

	fmt.Println("split a-b-c-d-e @ - ->>>>",strings.Split("a-b-c-d-e", "-"))

}