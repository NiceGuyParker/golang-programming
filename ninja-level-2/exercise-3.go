package main

import "fmt"

const (
	a int = 42
	b     = "James Bond"
	c     = 42.42
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T\t%T\t%T", a, b, c)
}
