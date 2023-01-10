package main

import (
	"fmt"
)

const boardSize int = 4 


func getBoard() []int {
	const boardVolume = boardSize * boardSize
	board := []int{}
	
	for i:=0; i < boardVolume; i++ {
		board = append(board, 0)
	}
	
	return board
}

func printBoard(board []int) {
	line := []int{}
	
	for i, elem := range board {
		line = append(line, elem)
		
		if i%boardSize == boardSize-1{
			fmt.Println(line)
			line = []int{}
		}
	}
}

func getInput() {
	var keyPressed string
	
	_, err := fmt.Scan(&keyPressed)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(keyPressed)
	}
}

func main() {
	// Variables
	board := getBoard()
	highScore := 0
	
	// game loop
	for ;; {
		// game info
		fmt.Println("\n \n \n ")
		fmt.Println("2048")
		fmt.Println("Higest score: ", highScore)
		
		// game functions
		printBoard(board)
		getInput()
	}
}
