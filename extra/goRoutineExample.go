package main

import "fmt";

func wrapper(a int,c chan int){
	result := a+10;
	c <- result;
}

func main() {
	c:= make(chan int)
	go wrapper(17,c)
	x:= <-c
	fmt.Println(x)
}