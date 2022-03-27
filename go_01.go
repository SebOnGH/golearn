package main

import "fmt"

var (
	stringOne string = "Hello"
	stringTwo string = "world"
	numberOne int    = 10
)

var (
	counter = 1
)

func main() {
	// Const
	const myConst int = 34
	fmt.Printf("%v,%T", myConst, myConst)

	// Primitive : int
	//var a int8 = 127
	//var auint uint = 255
	b1 := 10
	b2 := 11
	c1 := 8 //2³

	// bitshift
	fmt.Println(c1 << 3)
	fmt.Println(c1 >> 3)

	fmt.Println(b1 + b2)
	fmt.Println(b1 / b2)
	fmt.Println(b1 % b2)
	//fmt.Printf("%v, %T\n", a, a)
	//fmt.Printf("%v, %T\n", auint, auint)

	// Primitive : boolean
	var n bool = true
	o := 1 == 1
	p := 1 == 2

	fmt.Printf("%v, %T\n", o, o)
	fmt.Printf("%v, %T\n", p, p)

	fmt.Printf("%v, %T\n", n, n)

	// variables
	//var i int = 10
	//var j float32 = 26.6
	//h := 20
	//println(i, h)

	fmt.Printf("%v, %T\n", stringOne, stringOne)

	// simple text
	println("Hello world!")

}
