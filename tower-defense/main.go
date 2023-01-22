package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Enemy struct {
	Position [2]int
	Health   int
	Steps    int
}

type Shoot struct {
	Shoot bool
	Tower [2]int
	Enemy [2]int
}

var MAX_DISTANCE int = 7

var PATH_POSITION = [...][2]int{{0, 0}, {2, 0}, {2, 2}, {4, 2}, {6, 2}, {6, 4}, {6, 6}, {8, 6}, {10, 6}, {12, 6}, {12, 8}, {12, 10}, {14, 10}, {14, 12}, {16, 12}, {18, 12}}
var TOWERS_POSITION = [...][2]int{{2, 6}, {10, 2}, {8, 10}, {16, 16}}


var PLACED_TOWERS = []int{}
var ENEMY_INFO = []Enemy{}

var TOWER = [...]string{"T", "T", "M", "M"}
var ENEMY = [...]string{"O", "O", "I", "I"}

var PATH = [...]string{"*", "*", "*", "*"}
var SPECIAL_PATH = [...]string{"+", "+", "+", "+"}

var BULLET string = "/"

var status string = "Playing"
var score int = 0
var towerHealt int = 250
var enemyDamage int = 25
var money int = 200

var isRaid bool = false

func getIndex(x, y, lineLength int) int {
	return y*lineLength + x
}

func getEven(n int) int {
	if n%2 != 0 {
		return n + 1
	} else {
		return n
	}
}

func contain(arr []int, elem int) bool {
	for _, elemArr := range arr {
		if elemArr == elem {
			return true
		}
	}
	
	return false
}

func delay(ms time.Duration) {
	time.Sleep(ms * time.Millisecond)
}

func generateBoard(size int) []string {
	board := []string{}
	
	evenSize := getEven(size)
	fullSquare := evenSize * evenSize
	
	for i := 0; i < fullSquare; i++ {
		// use "-" only for debugging, for playing use space (" ")
		board = append(board, " ")
	}
	
	return board
}

func printBoard(board []string) {
	lineLength := int(math.Sqrt(float64(len(board))))
	
	fmt.Printf("\n")
	fmt.Println("Game status: ", status)
	fmt.Println("Score: ", score)
	fmt.Println("Main tower health: ", towerHealt)
	fmt.Println("Money: ", money)
	fmt.Printf("\n")
	for i, elem := range board {
		fmt.Printf(" %s", elem)
		
		if i%lineLength == lineLength-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func appendBlock(x, y, i, lineLength int, elem string, board []string) []string {
	newBoard := board
	
	switch i {
	case 0:
		newBoard[getIndex(x, y, lineLength)] = elem
	case 1:
		newBoard[getIndex(x+1, y, lineLength)] = elem
	case 2:
		newBoard[getIndex(x, y+1, lineLength)] = elem
	case 3:
		newBoard[getIndex(x+1, y+1, lineLength)] = elem
	}
	
	return newBoard
}

func appendPath(board []string) []string {
	newBoard := board
	lineLength := int(math.Sqrt(float64(len(board))))
	
	for pathIndex, xy := range PATH_POSITION {
		x := xy[0]
		y := xy[1]
		pathType := PATH
		
		if pathIndex == len(PATH_POSITION)-1 || pathIndex == 0 {
			pathType = SPECIAL_PATH
		}
		
		for i, elem := range pathType {
			newBoard = appendBlock(x, y, i, lineLength, elem, board)
		}
	}
	
	return newBoard
}

func appendEmptyTowers(board []string) []string {
	newBoard := board
	lineLength := int(math.Sqrt(float64(len(board))))
	
	for towerIndex, towerPosition := range TOWERS_POSITION {
		x := towerPosition[0]
		y := towerPosition[1]
		
		for towerSide := 0; towerSide < 4; towerSide++ {
			newBoard = appendBlock(x, y, towerSide, lineLength, strconv.Itoa(towerIndex), board)
		}
	}

	return newBoard
}

func appendTowers(board []string) []string {
	newBoard := board
	lineLength := int(math.Sqrt(float64(len(board))))
	
	for _, towerIndex := range PLACED_TOWERS {
		x := TOWERS_POSITION[towerIndex][0]
		y := TOWERS_POSITION[towerIndex][1]
		
		for index, elem := range TOWER {
			newBoard = appendBlock(x, y, index, lineLength, elem, board)
		}
	}

	return newBoard
}

func appendEnemy(board []string) []string {
	newBoard := board
	lineLength := int(math.Sqrt(float64(len(board))))
	
	for _, enemy := range ENEMY_INFO {
		x := enemy.Position[0]
		y := enemy.Position[1]
		
		for index, elem := range ENEMY {
			newBoard = appendBlock(x, y, index, lineLength, elem, board)
		}
	}
	
	return newBoard
}

func addEnemy() {
	var newEnemy Enemy
	newEnemy.Health = 75
	newEnemy.Steps = 0
	newEnemy.Position = PATH_POSITION[newEnemy.Steps]
	
	ENEMY_INFO = append(ENEMY_INFO, newEnemy)
}

func moveEnemy() {
	newEnemy := []Enemy{}
	
	for index := range ENEMY_INFO {
		// update enemy
		enemy := ENEMY_INFO[index]
		if enemy.Steps+1 < len(PATH_POSITION) {
			enemy.Steps += 1
		}
		enemy.Position = PATH_POSITION[enemy.Steps]
		
		// check if you are at the end of the path
		if enemy.Steps == len(PATH_POSITION)-1 {
			towerHealt -= enemyDamage 
		}
		newEnemy = append(newEnemy, enemy)
	}
	
	ENEMY_INFO = newEnemy
}

func getDistance(posTower, posEnemy [2]int) int {
	tX := posTower[0]
	tY := posTower[1]
	
	eX := posEnemy[0]
	eY := posEnemy[1]
	
	// round distance formula ( d = sqrt (x2-x1)**2 + (y2-y1)**2 )
	return int(math.Sqrt(float64((tX-eX)*(tX-eX) + (tY-eY)*(tY-eY))))
}

func sign(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	} else { 
		return 0
	}
}

// Bresenham's algorithm
func drawLine(x1, y1, x2, y2 int, board []string) []string {
	newBoard := board
	lineLength := int(math.Sqrt(float64(len(board))))
	
	x := x1
	y := y1
	
	dx := math.Abs( float64(x2-x1) )
	dy := math.Abs( float64(y2-y1) )
	
	s1 := sign(x2-x1)
	s2 := sign(y2-y1)
	
	var isChanged bool
	
	if dy > dx {
		temp := dx
		
		dx = dy
		dy = temp
		isChanged = true
	} else {
		isChanged = false
	}
	
	err := 2*dy - dx
	a := 2*dy
	b := 2*dx
	
	newBoard[getIndex(x,y,lineLength)] = BULLET
	
	for i:=0; i < int(dx); i++ {
		if err < 0 {
			err += a
			
			if isChanged {
				y += s2
			} else {
				x += s1
			}
		} else {
			y += s2
			x += s1
			err += b
		}
		
		newBoard[getIndex(x,y,lineLength)] = BULLET
	}
	
	return newBoard
}

func shootBullet(towerPosition, enemyPosition [2]int, board []string) []string {
	towerX := towerPosition[0]
	towerY := towerPosition[1]
	
	enemyX := enemyPosition[0]
	enemyY := enemyPosition[1]
	
	return drawLine(towerX, towerY, enemyX, enemyY, board)
}

func killEnemy(enemy Enemy, index int) {
	newEnemy := []Enemy{}
	
	for i,enemyObj := range ENEMY_INFO {
		if i != index {
			newEnemy = append(newEnemy, enemyObj)
		} else {
			var tempEnemy Enemy
			tempEnemy.Health = enemy.Health-20
			tempEnemy.Position = enemy.Position
			tempEnemy.Steps = enemy.Steps
			
			if tempEnemy.Health > 0 {
				newEnemy = append(newEnemy, tempEnemy)
			} else {
				score += 20
				money += 50
				towerHealt += 20
			}
		}
	}
	
	ENEMY_INFO = newEnemy
}

func shootTower() []Shoot {
	stages := []Shoot{}
	
	for _, towerIndex := range PLACED_TOWERS {
		tower := TOWERS_POSITION[towerIndex]
		
		for enemyIndex, enemy := range ENEMY_INFO {
			distance := getDistance(tower, enemy.Position)
			
			if distance <= MAX_DISTANCE {
				var stage Shoot 
				stage.Shoot = true
				stage.Tower = tower
				stage.Enemy = enemy.Position
				
				stages = append(stages, stage)
				
				killEnemy(enemy, enemyIndex)
			}
		}
	}
	
	return stages
}

func raid() {
	if isRaid {
		addEnemy()
	}
}

func main() {
	inputChannel := make(chan int)
	
	// intput
	go func() {
		for ;; {
			input := -1
			fmt.Scanf("%d", &input)
			
			inputChannel <- input
		}
	}()
	
	go func() {
		// game loop
		for ;; {
			board := generateBoard(20)
			board = appendPath(board)
			board = appendEmptyTowers(board)
			board = appendTowers(board)
			
			moveEnemy()
			raid()
			
			board = appendEnemy(board)
			
			// shoot bullet
			if len(shootTower()) != 0 {
				for _, shoot := range shootTower() {
					board = shootBullet( shoot.Tower, shoot.Enemy, board )
				}
			} 
			
			// game status
			if	towerHealt <= 0 {
				status = "Lost"
			}
			
			// raid
			if len(ENEMY_INFO) <= 6 {
				isRaid = true 
			} else {
				isRaid = false
			}
			
			printBoard(board)
			delay(1250)
			
			// check for input
			select {
				case input := <- inputChannel:
					if input != -1 && !contain(PLACED_TOWERS, input) && money >= 150 {
						money -= 150
						PLACED_TOWERS = append(PLACED_TOWERS, input)
					}
				default: // this should never happend
			}
		}
	}()
	
	// check for end?
	select {}
}
