package main

import "fmt"

func main() {
	a := 42 == 42
	b := 100 <= 42
	c := 200 >= 42
	d := 72 != 72
	e := 55 > 45
	f := 21 < 45

	fmt.Println(a, b, c, d, e, f)

}
