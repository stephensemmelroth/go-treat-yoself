package main

import "fmt"

func main() {
	var counting_numbers [9]int // make an empty array
	
	fmt.Println("Empty array:", counting_numbers)
	counting_numbers = [9]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(counting_numbers)
	
	var a [3]int = [3] int {4,5,6}  // White space, don't care
	fmt.Println(a)
	
	ctn_nums :=	[9]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(ctn_nums)
	
	fmt.Println("The third counting number:", counting_numbers[2])
	fmt.Println("First three counting numbers:", counting_numbers[:3]) // Doesn't include end of slice range
	fmt.Println("Not first three counting numbers:", counting_numbers[3:]) // Includes beggining of slice
	fmt.Println("Index 3-5?", counting_numbers[3:5])  // includes start but not finish (else would be 4,5,6)
	fmt.Printf("[3:5] is a %T (a slice).\n", counting_numbers[3:5])
	
	var my_slice []int = []int{1,2,3} // Declare a slice
	fmt.Println("My slice", my_slice)
	
	not_string := []rune{'n','o','t','s','t','r'}
	fmt.Printf("Is it a string?: %T\n", not_string)
	fmt.Printf("Is it a string?: %T\n", "A string")
	
	
}