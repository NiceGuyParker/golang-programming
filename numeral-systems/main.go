package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	//converting s to a slice of byte
	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	//percent T print type of n
	fmt.Printf("%T\n", n)
	//percent b prints n in binary
	fmt.Printf("%b\n", n)
	//percent x prints n in hexadecimal. adding the hash symbol prints in hexadecimal with 0X in front
	fmt.Printf("%#X\n", n)
}
