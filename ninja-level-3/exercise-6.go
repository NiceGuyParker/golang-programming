package main

import "fmt"

func main(){
	x := 0
	for {
		if x > 1000{
			break
		}
		fmt.Println(x, x/2)
		x++
	}

}
