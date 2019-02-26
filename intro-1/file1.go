package main

import ("fmt"
		"math"
		"math/rand")

func foo() {
	//capital indicates that the method is imported
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func add(x float64, y float64) float64 {
//if same types can say
// func add(x,y float64) float65
	return x+y;;
}

func numberOneToN() {
	/*
		godoc math/rand Intn
	*/
	fmt.Println("A number from 1-100", rand.Intn(100))
}

func multipleReturns(a,b string) (string,string) {
	return a,b
}

func main() {
	foo()
	numberOneToN()

	//define variables
	var num1 float64 = 5.6
	var num2 float64 = 9.5

	//OR can say 
	var num3, num4 float64 = 5.6, 9.5

	//OR within funcs only:
	num5,num6 := 5.6, 9.5

	//constants
	const x int = 5

	//type casting
	var a int = 64
	var b float64 = float64(a)

	//auto set var
	x := a //x will be int
	fmt.Println(add(num1,num2))

}