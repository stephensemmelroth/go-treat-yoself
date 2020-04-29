package main

import . "fmt"

func main() {
	counter := []int{4,5,6,7,8}
	Println(counter)
	
	for i, e := range counter {
		Println("Index", i, "has element", e)
	}
	
	for _, e := range counter {
		Println(e)
	}
	for i, _ := range counter {
		Println(i)
	}
	
	dict := map[string]string{
		"Tired" : "The daily state of a programmer",
 		"Code": "The product of a programmer", 
 		"Coffee" : "Transforms into code", 
 		}
		
	for k, v := range dict {
		Println("The key is", k, "and value is", v)
	}
	
	for k := range dict {
		Println("Key", k)
	}
	
	str := "I'm a fuzzy unicorn!"
	for i, r := range str{
		Println("Byte index:", i, "rune (NOT BYTE):", r)
	}
}