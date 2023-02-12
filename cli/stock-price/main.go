package main

import (
	"fmt"
	"sort"
	"strconv"
)


const SEPARATE string = "Ôö╝"
const SEPARATE_UP string = "ÔöČ"
const SEPARATE_DOWN string = "Ôö┤"

const TURN_LEFT_UP string = "ÔĽ«"
const TURN_RIGHT_UP string = "ÔĽş"
const TURN_LEFT_DOWN string = "ÔĽ»"
const TURN_RIGHT_DOWN string = "ÔĽ░"
const LINE_COLUMN string = "Ôöé"
const LINE_ROW string = "ÔöÇ"

const POINT string = "┬Ą"

func get_lowest_value(arr []int) int {
	sort.Ints(arr)

	return arr[0]
}

func get_higest_value(arr []int) int {
	sort.Ints(arr)

	return arr[len(arr)-1]
}

// return bigest length of numbers in array
func get_max_len(arr []int) int {
	var maxNumber int

	// loop trough arr and check if is anything bigger than maxNumber
	for _, value := range arr {
		str := strconv.Itoa(value)

		if len(str) > maxNumber {
			maxNumber = len(str)
		}
	}

	return maxNumber
}

// get index of element in array
func indexOf(arr []int, num int) int {
	for index, elem := range arr {
		if elem == num {
			return index
		}
	}

	return -1
}

// get spearating symbol based on index (for better look)
func get_separate_symbol(arr []int, index int) string {
	if index == 0 {
		return SEPARATE_UP
	}

	if index == len(arr)-1 {
		return SEPARATE_DOWN
	}

	return SEPARATE
}
/*

How this will work?
I don't know too..

get range of price and store it in array
	[0] MAX = MAX + MIN

	get DIFFerence between ( MAX/MIN )
	fill range arr with DIFF * ( MAX - MIN )? prob.
		if is something really close to it, change it with price value
		else add it to array

	for each RANGE ARRAY
		if elem is in VALUE print ( index in array * spaces )
		else print ELEM //TODO

	[N] MIN = MIN - MIN
*/


func get_space(x int) string {
	space := ""
	for i:=0; i < x; i++ {
		space = space + " "
	}

	return space
}

// return spaces depeding on length number
func generate_space(max_num, current int) string {
	var space string
	str := strconv.Itoa(current)

	if len(str) < max_num {
		space = get_space(max_num-len(str))
	}

	return space
}


// return range from higest -> lowest
func get_range(min, max int) []int {
	arr := []int{}

	for i:=max; i >= min; i-- {
		arr = append(arr, i)
	}

	return arr
}

// print chart
func cprint(sorted []int, main []int) {
	higestValue := sorted[len(sorted)-1]
	lowestValue := sorted[0]
	numberRange := get_range(lowestValue, higestValue)

	for index, value := range numberRange {
		var resultRow string

		// get spaces and format number + |
		spaces := generate_space(get_max_len(sorted), value)
		formatedRange := fmt.Sprintf("%s%d %s", spaces, value, get_separate_symbol(numberRange, index))

		// elements
		var formatedItems string
		if indexOf(main, value) != -1 {
			// +1 for better look
			spaces := get_space(indexOf(main,value)+1)
			formatedItems = fmt.Sprintf("%s%s", spaces, POINT)
		}

		// connect print row
		resultRow = fmt.Sprintf("%s%s", formatedRange, formatedItems)

		fmt.Println(resultRow)
	}
}

func main() {
	fmt.Println("Get stock price for BTC")

	// get currency and sort it
	items := []int{ 3,7,9,15,8,11,4,14 }
	var currency []int
	var sortedCurrency []int

	// duplicate
   currency = append(currency, items...)
   sortedCurrency = append(sortedCurrency, items...)

	// sort currency
	sort.Ints(sortedCurrency)

	// print chart
	cprint(sortedCurrency, currency)
}
