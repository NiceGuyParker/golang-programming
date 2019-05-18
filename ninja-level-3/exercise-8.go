package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("will not print")

	case (2 == 2):
		fmt.Println("this will print")
	}
}
