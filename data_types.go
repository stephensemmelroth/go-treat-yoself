package main

import "fmt"

func main() {
	var x int
	x = 10
	
	fmt.Println(x, "is an int, lol")
	
	var a, b, c = 1, 2, 3
	fmt.Println(a,b,c)
	
	name := (a + b + c)
	fmt.Println(name)
	
	pi := 3.14; fmt.Printf("Pi is a %T.\n", pi)
	
	const pi_2 float32 = 3.1415926
	fmt.Printf("Pi_2 is a %T.\n", pi_2)
	
	var (
		r int = 10
		pi_3 float32 = 3.1415
	)
	fmt.Println("AREA!",float32(r*r)*pi_3)
	
	var y int8 = 5
	fmt.Printf("x is a %T and y is a %T.\n", x, y)
	// z := x + y will result in an invalid operation (adding int and int 8)
	z := int(x) + int(y)
	fmt.Printf("z is a %T.\n", z)
	
}