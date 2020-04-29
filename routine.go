package main

import . "fmt"

func power_up(p int, c chan int){
	// Prints here may not make it to the terminal before termination
	Println(`
(39 episodes later...)
HAAAAAAA!!!!!!
HAAAAAAAAAAAAAAAAA!!!!!!!!!!!!!!
KAIOKEN TIMES TEEEEENNNNNNN!!!!!`)
	c <- (p * 10)
}

func crush(c chan int){
	close(c)
}

func main() {
	// Goku, without kaioken
	son_goku := 901
	Println("Current power:", son_goku)
	
	scouter := make(chan int)
	go power_up(son_goku, scouter)
	
	prince_vegeta := <- scouter
	
	if prince_vegeta > 9000 {
		Println("\nNo, it isn't broken...\nHis power level's OVER NINE THOUSAAAND!")
	}
	
	crush(scouter)
}