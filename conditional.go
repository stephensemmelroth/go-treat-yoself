package main

import . "fmt"

func main() {
	i := 5
	if i < 5 {
		Println("LT 5!")
	} else if i == 1 {
		Println("EQ 1!")
	} else {
		Println("Blarg")
	}
	
	x, y := 5, 6
	if z := x*y % 2; z == 0 {
		Println("Z is even!", z)
	} else {
		Println("Z is odd!", z)
	}
	
	a := 5
	if a == 5 {
		b := 10
		Println(b)
	} else {
		// Reference here would produce compile time error
		//   EVEN IF reaching here is impossible
		// Println(b)
		Println("These aren't the droids we're looking for")
	}
	
	var check interface{}  // Must have an interface for type assertions
	check = "String?!"
	if str, ok := check.(string); ok {
		Printf("I'M A %T.\n", str)
	}
	
	letter := "a"
	switch letter {
		case "a":
			Println(letter)
			fallthrough 
		case "b":
			Println("Get here 1?")
			break
		default:
			Println("Get here 2?")
	}
	
	number := 5
	switch {
		case number == 5: 
			Println("FIVE")
		case number > 5:
			Println("GT FIVE!")
		case number < 5:
			Println("LT FIVE!")
		default:
			Println("Blarg")
	}
}