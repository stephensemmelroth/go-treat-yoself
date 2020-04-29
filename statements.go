package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("The type of a is %T\n", a)
	
	if b := 10; b > 5 {
		fmt.Println(true)
	}
	
	s := `This is a 
multi-line string`
	fmt.Println(s)
	
	c := 1 + 2 +
	3 + 4 + 5
	
	fmt.Println(c)
}