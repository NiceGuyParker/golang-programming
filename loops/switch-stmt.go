package main

import "fmt"

func main(){
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print 2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
		//prints the below statement only with fallthrough keyword
	case (4 == 4):
		fmt.Println("also true, doesnt print")

	}
}
