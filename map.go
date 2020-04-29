package main

import . "fmt"

func main() {
	var my_map map[string]int
	my_map = make(map[string]int)
	Println(my_map)
	
	var map_2 map[string]string
	map_2 = map[string]string{"Key_1": "value_1", "K_2":"V_2"}
	Println(map_2)
	
	dict := map[string]string{
		"Tired" : "The daily state of a programmer",
 		"Code": "The product of a programmer", 
 		"Coffee" : "Transforms into code", 
 		}
	Println(dict)
	
	var dict_2  map[string]string = map[string]string{
		"Tired" : "The daily state of a programmer",
 		"Code": "The product of a programmer", 
 		"Coffee" : "Transforms into code", 
 		}
	Println(dict_2)
	
	// Assign directly into the mapping
	my_map["Scouter's Broken"] = 9001
	Println(my_map["Scouter's Broken"])
	// Delete a mapping
	delete(my_map, "Scouter's Broken")
	// Use the , OK idiom to check if it's gone
	value, okay := my_map["Scouter's Broken"]
	Println(value, okay)
}