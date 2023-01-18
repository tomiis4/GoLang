package main

import (
	"fmt"
	"math"
)


var PATH_POSITION= [...][2]int{ {0,0}, {2,0}, {2,2}, {4,2}, {6,2}, {6,4}, {6,6}, {8,6}, {10,6} }
var TOWER = [...]string{"T","T","M", "M"}
var ENEMY = [...]string{"O","O","I", "I"}
// var PATH = [...]string{"_","|","|", "_"}
var PATH = [...]string{"*","*","*", "*"}

func getIndex(x int, y int, lineLength int) int {
	return y * lineLength + x
}

func generateBoard(size int) []string {
	board := []string{}
	fullSquare := size*size
	
	for i:=0; i < fullSquare; i++ {
		// use "-" only for debugging
		board = append(board, "-")
	}
	
	return board
}

func printBoard(board []string) {
	lineLength := int( math.Sqrt( float64(len(board)) ) )
	
	fmt.Printf("\n")
	for i, elem := range board {
		fmt.Printf(" %s", elem)
		
		if i%lineLength == lineLength-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func appendPath(board []string) []string {
	newBoard := board
	lineLength:= int( math.Sqrt( float64(len(board)) ) )
	
	for _, xy := range PATH_POSITION {
		x := xy[0]
		y := xy[1]
		
		for i, elem := range PATH {
			switch i {
				case 0:
					newBoard[ getIndex(x,y, lineLength) ] = elem
				case 1:
					newBoard[ getIndex(x+1,y, lineLength) ] = elem
				case 2:
					newBoard[ getIndex(x,y+1, lineLength) ] = elem
				case 3:
					newBoard[ getIndex(x+1,y+1, lineLength) ] = elem
			}
		}
	}
	
	return newBoard
}

func main() {
	board := generateBoard(15)
	board = appendPath(board)
	
	printBoard(board)
}
