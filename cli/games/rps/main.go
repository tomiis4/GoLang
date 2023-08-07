package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randInt(n int) int {
	newTime := rand.NewSource(time.Now().UnixNano())
	resetRandom := rand.New(newTime)
	
	randomNumber := resetRandom.Intn(n)
	
	return randomNumber
}

func getStatus(plr, pc string) {
	// plr win
	if plr == "ROCK" && pc == "SCISSORS" {
		fmt.Printf("Player has won, you picked %s and computer picked %s", plr, pc)
	} else if plr == "PAPER" && pc == "ROCK" {
		fmt.Printf("Player has won, you picked %s and computer picked %s", plr, pc)
	} else if plr == "SCISSORS" && pc == "PAPER" {
		fmt.Printf("Player has won, you picked %s and computer picked %s", plr, pc)
	} 
	
	// lose
	if plr == "ROCK" && pc == "PAPER" {
		fmt.Printf("Player has lost, you picked %s and computer picked %s", plr, pc)
	} else if plr == "PAPER" && pc == "SCISSORS" {
		fmt.Printf("Player has lost, you picked %s and computer picked %s", plr, pc)
	} else if plr == "SCISSORS" && pc == "ROCK" {
		fmt.Printf("Player has lost, you picked %s and computer picked %s", plr, pc)
	}
	
	// tie
	if plr == pc {
		fmt.Printf("No one won, you picked %s and computer picked %s", plr, pc)
	}
	
	fmt.Printf("\n\n")
}

func main() {
	items := [...]string{"ROCK", "PAPER", "SCISSORS"}
	
	for ;; {
		// input
		var input int
		
		fmt.Println("Enter digits")
		fmt.Println("0=ROCK | 1=PAPER | 2=SCISSORS")
		fmt.Scanf("%d", &input)
		
		playerInput := items[input]
		
		// pc input
		pcItem := items[randInt(len(items))]
		
		getStatus(playerInput, pcItem)
	}
}
