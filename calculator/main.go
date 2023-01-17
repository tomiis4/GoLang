package main

import "fmt"

func calculate(operator string, first int, second int) int {
	switch operator {
		case "+":
			return first + second 
		case "-":
			return first - second 
		case "*":
			return first * second 
		case "/":
			return first / second 
		default: return -1
	}
}

func main() {
	var firstValue int
	var secondVaue int
	var operator string
	
	fmt.Printf("Enter first value: ")
	fmt.Scanf("%d\n", &firstValue)
	
	fmt.Printf("Enter second value: ")
	fmt.Scanf("%d\n", &secondVaue)
	
	fmt.Printf("Enter operator: ")
	fmt.Scanf("%s", &operator)
	
	fmt.Println("Result is:", calculate(operator, firstValue, secondVaue))
}
