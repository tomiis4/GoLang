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


// x = previsou number
// y = current number
// current = current index of line you are printing

/*

TODO



CURRENT MUST BE <= Y    >= X

moving from DOWN to UP (y > x)
	current == y write turing up right
	current == x write up/down
	else write up/down

moving from UP to DOWN (y < x)
	current == y write turn bottom right
	current == x write up/down
	else write up/down
*/

func get_symbol(x,y, current int) string {
	// move up
	isUp := y > x
	canWriteUp := current <= y && current >= x && isUp

	// move down
	isDown:= y < x
	canWriteDown := current >= y && current <= x && isDown

	// if current is at point
	if current == x || current == y {
		return POINT
	}

	if canWriteDown{
		return TURN_LEFT_DOWN
	}
	if canWriteUp {
		return TURN_RIGHT_DOWN
	}


	return " "
}

// functions return string with spaces
func get_space(n int) string {
	space := ""
	for i:=0; i < n; i++ {
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

func bottom_status(sorted []int) []string {
	full_chart := []string{} 

	higestValue := sorted[len(sorted)-1]
	lowestValue := sorted[0]
	numberRange := get_range(lowestValue, higestValue)

	for index, value := range numberRange {
		var resultRow string


		// get spaces and format number + |
		spaces := generate_space(get_max_len(sorted), value)
		formatedRange := fmt.Sprintf("%s%d %s", spaces, value, get_separate_symbol(numberRange, index))


		var formatedItems string


		// connect and add to full_chart
		resultRow = fmt.Sprintf("%s%s", formatedRange, formatedItems)
		full_chart = append(full_chart, resultRow)
	}

	return full_chart
}

// TODO make function for path creating
//	FIXME OFFTOPICS install "languages" for nvim

// Prob get one array with points and second with path, then connect which will add points to path :)
// print chart
func cprint(sorted []int, main []int) {
	chart_numbers := bottom_status(sorted)

	// get lowest number from range
	// numberRange := get_range(sorted[0], len(sorted)-1)
	// lowest_value := numberRange[len(numberRange)-1]

	// // add points
	// for index, elem := range main {
	// 	// display points
	// 	spaces := get_space(index+1)
	// 	line_index := len(chart_numbers)+lowest_value-elem-1


	// 	// this one work (add POINTs)
	// 	chart_numbers[line_index] = fmt.Sprintf("%s%s%s", chart_numbers[line_index], spaces, POINT)
	// }

	// loop column by column
	for x:=0; x < len(main)-1; x++ {
		for y:=0; y < len(chart_numbers); y++ {
			// letter := main[x]
			// future_letter := main[x+1]
			// spaces := get_space(x)

			// symbol := get_symbol(letter, future_letter, y)

			// fmt.Println(spaces, letter, future_letter)
			// chart_numbers[x] = fmt.Sprintf("%s%s%s", chart_numbers[x], "", symbol)
			fmt.Println(main[x])
			chart_numbers[y] = fmt.Sprintf("%s#", chart_numbers[y])
		}
	}

	// add path
	// for each letter it it loops column in chart_numbers
			// if column+1 is in main
			// letter := main[column]
			// future_letter := main[column+1]
			// fmt.Println(letter, future_letter,i)
			// fmt.Println(column,i)
			// if index+1 exists
			// if len(main) > index+1 {
				// spaces := get_space(index+1)
				// symbol := "#"
				// symbol := get_symbol(elem, main[index+1], line_index)
				// chart_numbers[index] = fmt.Sprintf("%s%s%s", chart_numbers[index], spaces, symbol)
			// }

	// print array
	for _, elem := range chart_numbers {
		fmt.Println(elem)
	}
}

func main() {
	fmt.Printf("Get stock price for BTC\n\n")

	// get currency and sort it
	items := []int{ 3,7,9,15,8,11,4,14,3 }
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
