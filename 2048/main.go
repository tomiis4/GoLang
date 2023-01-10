package main

import (
	"fmt"
	
	"math/rand"
	"time"
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

func randomNum(board []int) []int {
	newBoard := board
	
	for i:=0; i < len(board)*2; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(board))
		
		if board[n] == 0 {
			newBoard[n] = 2
			return newBoard
		}
	}
	
	return []int{0} 
}

func mergeLine(line []int) []int {
	newLine := line
	
	for i:=0; i < len(newLine)-1; i++ {
		if newLine[i] == newLine[i+1] {
			newLine[i] = newLine[i]*2
		}
	}
	
	return newLine
}

func shiftX(board []int, side string) []int {
	newBoard := []int{} 
	line := []int{}
	
	for i, elem := range board {
		line = append(line, elem)
		
		// for each line
		if i%boardSize == boardSize-1{
			// shift left
			if side == "left" {
				j := 0
				for k:=0; k < len(line); k++ {
					if line[k] != 0 {
						line[j] = line[k]
						j++
					}
				}
				
				// fill with 0
				for k := j; k < len(line); k++ {
				   line[k] = 0
				}
			}
			
			// shift to right 
			if side == "right" {
				j := len(line)-1
				for k:=len(line)-1; k >= 0; k-- {
					if line[k] != 0 {
						line[j] = line[k]
						j--
					}
				}
				
				// fill with 0
				for k := 0; k <= j; k++ {
				   line[k] = 0
				}
			}
			
			line = mergeLine(line)
			
			// append to array
			for k:=0; k < len(line); k++ {
				letter := line[k]
				
				newBoard = append(newBoard, letter)
			}
			
			line = []int{}
		}
	}
	
	return newBoard
}

func getInput(board []int) []int {
	var keyPressed string
	var newBoard = board
	
	_, err := fmt.Scan(&keyPressed)
	
	if err != nil {
		panic(err)
	}
	
	switch keyPressed {
		case "w":
			randomNum(newBoard)
		case "s":
			randomNum(newBoard)
		case "a":
			newBoard = shiftX(board, "left")
			randomNum(newBoard)
		case "d":
			newBoard = shiftX(board, "right")
			randomNum(newBoard)
	}
	
	return newBoard
}

func main() {
	// Variables
	board := getBoard()
	highScore := 0
	
	// game loop
	for ;; {
		// game info
		fmt.Println("\n \n ")
		fmt.Println("2048")
		fmt.Println("Higest score: ", highScore)
		
		// game functions
		printBoard(board)
		board = getInput(board)
	}
}
