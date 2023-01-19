package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Enemy struct {
	Position [2]int
	Health int
	Steps int
}

var PATH_POSITION = [...][2]int{ {0,0}, {2,0}, {2,2}, {4,2}, {6,2}, {6,4}, {6,6}, {8,6}, {10,6} }

var PLACED_TOWERS = []int{ 2 }
var TOWERS_POSITION = [...][2]int{ {0,4}, {10,2}, {4,8} }

var ENEMY_INFO = []Enemy{}

var TOWER = [...]string{"T","T","M", "M"}
var ENEMY = [...]string{"O","O","I", "I"}

var PATH = [...]string{"*","*","*", "*"}
// var PATH = [...]string{"_","|","|", "_"}


func getIndex(x int, y int, lineLength int) int {
	return y * lineLength + x
}

func getEven(n int) int {
	if n%2 != 0 {
		return n+1
	} else {
		return n
	}
}

func delay(ms time.Duration) {
	time.Sleep(ms * time.Millisecond)
}

func generateBoard(size int) []string {
	board := []string{}
	
	evenSize := getEven(size)
	fullSquare := evenSize*evenSize
	
	for i:=0; i < fullSquare; i++ {
		// use "-" only for debugging, for playing use space (" ")
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

func appendBlock(x int, y int, i int, lineLength int, elem string, board []string) []string {
	newBoard := board
	
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
	
	return newBoard
}

func appendPath(board []string) []string {
	newBoard := board
	lineLength:= int( math.Sqrt( float64(len(board)) ) )
	
	for _, xy := range PATH_POSITION {
		x := xy[0]
		y := xy[1]
		
		for i, elem := range PATH {
			newBoard = appendBlock(x,y, i, lineLength, elem, board)
		}
	}
	
	return newBoard
}

func appendEmptyTowers(board []string) []string {
	newBoard := board
	lineLength := int( math.Sqrt( float64(len(board)) ) )
	
	for towerIndex, towerPosition := range TOWERS_POSITION {
		x := towerPosition[0]
		y := towerPosition[1]
		
		for towerSide:=0; towerSide < 4; towerSide++ {
			newBoard = appendBlock(x, y, towerSide, lineLength, strconv.Itoa(towerIndex), board)
		}
	}
	
	return newBoard
}

func appendTowers(board []string) []string {
	newBoard := board
	lineLength := int( math.Sqrt( float64(len(board)) ) )
	
	for _, towerIndex := range PLACED_TOWERS {
		x := TOWERS_POSITION[towerIndex][0]
		y := TOWERS_POSITION[towerIndex][1]
		
		for index, elem := range TOWER {
			newBoard = appendBlock(x,y, index, lineLength, elem, board)
		}
	}
	
	return newBoard
}

func appendEnemy(board []string) []string {
	newBoard := board
	lineLength := int( math.Sqrt( float64(len(board)) ) )
	
	for _, enemy := range ENEMY_INFO {
		x := enemy.Position[0]
		y := enemy.Position[1]
		
		for index, elem := range ENEMY {
			newBoard = appendBlock(x,y, index, lineLength, elem, board)
		}
	}
	
	return newBoard
}

func addEnemy() {
	var newEnemy Enemy
	newEnemy.Health = 100
	newEnemy.Steps = 0
	newEnemy.Position = PATH_POSITION[newEnemy.Steps]
	
	ENEMY_INFO = append(ENEMY_INFO, newEnemy)
}

func moveEnemy() {
	newEnemy := []Enemy{}
	
	for index:= range ENEMY_INFO {
		// update enemy
		enemy := ENEMY_INFO[index]
		if enemy.Steps+1 < len(PATH_POSITION) {
			enemy.Steps += 1
		}
		enemy.Position = PATH_POSITION[enemy.Steps]
		
		// check if you are at the end of the path
		if enemy.Steps == len(PATH_POSITION)-1 {
			//TODO add remove health from main "tower"
		} else {
			newEnemy = append(newEnemy, enemy)
		}
	}
	
	ENEMY_INFO = newEnemy
}

func main() {
	// testing only
	addEnemy()
	
	// game loop
	for {
		board := generateBoard(15)
		board = appendPath(board)
		board = appendEmptyTowers(board)
		board = appendTowers(board)
		
		moveEnemy()
		
		board = appendEnemy(board)
		
		printBoard(board)
		delay(1250)
	}
}
