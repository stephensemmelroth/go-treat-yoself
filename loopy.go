package main

import . "fmt"

func main() {
	for index := 0; index < 10; index++ {
		Println("I'm a classic for loop!", index)
	}
	
	i := 0
	for i < 5 {
		Println(i)
		i++
	}
	
	idx := true
	for idx {
		Println("I'm in danger!")
		idx = false
	}
	
	i = 0
	for {  // This isn't watching for ANY exit condition implicitly
		Println("MWAHAHAHAHA")
		if !idx {
			break  // MR WIZARD MR WIZARD!
		}
	}
	
	stop := false
	for { 
		if stop {
			break
		}
		Println("ZOOM!")
		stop = !stop
	}
}