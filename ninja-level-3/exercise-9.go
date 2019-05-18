package main

import "fmt"

func main() {
	favSport := "basketball"
	switch favSport {
	case "skiing":
		fmt.Println("hit that black diamond!")
	case "soccer":
		fmt.Println("it's football man!")
	case "basketball":
		fmt.Println("I just posterized you!")
	}
}
