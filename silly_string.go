package main

import . "fmt"

func main() {
	var my_string string = "This is a string with double quotes"
	str_2 := "Another string"
	str_3 := "\x6e\x6f\x74\x20\x75\x74\x66\x2d\x38"  // "not utf-8"
	
	Println(str_3)
	
	Printf("%x\n", my_string)
	
	Println(str_2[0], "is a RUNE!!!!! X_X")
	Println(str_2[1:4])
	Println(str_2[8:])
}