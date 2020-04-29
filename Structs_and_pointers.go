package main

import . "fmt"


type coord struct{
	x int
	y int
}

func (c coord) smoosh() int {
	return (c.x + c.y)
}	

// Mutator
func (c *coord) slope(b int) {  // Returns nothing, mutates
	c.y = c.x + b
}

func main() {
	var p *int = new(int)
	
	a := 10
	q := &a
	
	Printf("%T.\n", p)
	Printf("%T.\n", q)
	
	var my_coord1 coord = coord{42, 42}
	Printf("%T.\n", my_coord1)
	
	my_coord := struct{ x, y int}{5, 5} // Anonymous structure
	Printf("%T.\n", my_coord)
	
	Println("The answer to life is", my_coord1.x)
	Println("The answer to life is", my_coord1.smoosh()/2)
	my_coord1.slope(9000)
	Println("His power level is", my_coord1.y)
	
}