package main

import (
	"fmt"
	"strconv"
)

func toInt(str *string) *int {
	value, _ := strconv.Atoi(*str)

	return &value
}

func main() {
	// create new pointer and allocate memory
	intPointer := new(int)
	// get value of address using *
	*intPointer = 5

	fmt.Println(*intPointer)


	// create new variable
	number := 9
	// read address from number
	numerPointer := &number
	fmt.Println(*numerPointer)

	str := "-1"
	num := toInt(&str)
	fmt.Println( *num )
}
