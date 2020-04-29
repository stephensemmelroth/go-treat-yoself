package main

import . "fmt"

func main() {
	var yen rune = '¥'
	char_2 := 'A'
	
	Println("Prints code values:", yen, char_2)
	Printf("Types: %T, %T.\n", yen, char_2)
	Printf("AS CHARACTERS! %q and %q.\n", yen, char_2)
	
	char_array := [...]rune{'b','o','b'}
	var char_slice []rune = []rune{'j','a','n','e'}
	Printf("%T vs %T.\n", char_array, char_slice)
}