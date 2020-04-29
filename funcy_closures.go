package main

import . "fmt"

	
func smoosh(v_1 int, v_2 float64) int {
	return v_1 + int(v_2)
}

func smooth(v int, x float64) (int, int) {
	return int(v), int(x)
}

//Closure/generator
func gen_int() func() int {
	i := 0
	return func() int { i++; return i }
}
	
func main() {
	Println(smoosh(7, 3.14))
	Println(smooth(7, 3.14))
	
	nextint := gen_int()
	i := 0 
	for i < 10 {
	   Println(i)
	   i = nextint()
	}
}