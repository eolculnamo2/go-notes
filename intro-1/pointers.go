package main

import "fmt"

func main() {
	x := 15
	//& points to something
	//in memory
	a := &x //memory address
	*a = 5
	fmt.Println(a) // prints memory address 0xc000008600 or something
	fmt.Println(*a) // represent value at memory address
	fmt.Println(x) // x is now 5 because they represent the same memory address
}