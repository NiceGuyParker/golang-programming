package main

import "fmt"

var x bool

func main(){
	a := 42
	b := 42
	fmt.Println(a >= b)
	fmt.Println(x)
	x = true
	fmt.Println(x)

}
