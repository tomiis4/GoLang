package main

import (
	"fmt"
	"sort"
	"strconv"
)

const SEPARATE string = "┼"
const SEPARATE_UP string = "┬"
const SEPARATE_DOWN string = "┴"

const TURN_LEFT_UP string = "╮"
const TURN_RIGHT_UP string = "╭"
const TURN_LEFT_DOWN string = "╯"
const TURN_RIGHT_DOWN string = "╰"
const LINE_COLUMN string = "│"
const LINE_ROW string = "─"

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
func get_separate_symbol(main, range_arr []int, number, index int) string {
	first_elem := main[0]
	is_line_start := first_elem == number

	// make long line
	if index == 0 && is_line_start {
		return fmt.Sprintf("%s%s", SEPARATE_UP, LINE_ROW)
	}

	if index == len(range_arr)-1 && is_line_start {
		return fmt.Sprintf("%s%s", SEPARATE_DOWN, LINE_ROW)
	}

	if is_line_start {
		return fmt.Sprintf("%s%s", SEPARATE, LINE_ROW)
	}

	// add separating
	if index == 0 {
		return fmt.Sprintf("%s ", SEPARATE_UP)
	}

	if index == len(range_arr)-1 {
		return fmt.Sprintf("%s ", SEPARATE_DOWN)
	}

	return fmt.Sprintf("%s ", SEPARATE)
}

// x = previsou number
// y = current number
// current = current index of line you are printing
func get_symbol(x, y, current int) string {
	// move up
	isUp := y > x
	canWriteUp := current <= y && current >= x && isUp

	// move down
	isDown := y < x
	canWriteDown := current >= y && current <= x && isDown

	// if they are at same line
	if x == y && current == x {
		return LINE_ROW
	}

	// if current is at point
	if current == x && isUp {
		return TURN_LEFT_DOWN
	}

	if current == x && isDown {
		return TURN_LEFT_UP
	}

	if canWriteDown && current == y {
		return TURN_RIGHT_DOWN
	}

	if canWriteUp && current == y {
		return TURN_RIGHT_UP
	}

	if canWriteDown {
		return LINE_COLUMN
	}
	if canWriteUp {
		return LINE_COLUMN
	}

	return ""
}

// functions return string with spaces
func get_space(n int) string {
	space := ""
	for i := 0; i < n; i++ {
		space = space + " "
	}

	return space
}

// return spaces depeding on length number
func generate_space(max_num, current int) string {
	var space string
	str := strconv.Itoa(current)

	if len(str) < max_num {
		space = get_space(max_num - len(str))
	}

	return space
}

// return range from higest -> lowest
func get_range(min, max int) []int {
	arr := []int{}

	for i := max; i >= min; i-- {
		arr = append(arr, i)
	}

	return arr
}

func bottom_status(sorted, main []int) []string {
	full_chart := []string{}

	higestValue := sorted[len(sorted)-1]
	lowestValue := sorted[0]
	numberRange := get_range(lowestValue, higestValue)

	for index, value := range numberRange {
		var resultRow string

		// get spaces and format number + |
		spaces := generate_space(get_max_len(sorted), value)
		formatedRange := fmt.Sprintf("%s%d %s", spaces, value, get_separate_symbol(main, numberRange, value, index))

		var formatedItems string

		// connect and add to full_chart
		resultRow = fmt.Sprintf("   %s%s", formatedRange, formatedItems)
		full_chart = append(full_chart, resultRow)
	}

	return full_chart
}

// print chart
func cprint(sorted []int, main []int) {
	chart_numbers := bottom_status(sorted, main)

	// get lowest number from range
	numberRange := get_range(sorted[0], len(sorted)-1)
	lowest_value := numberRange[len(numberRange)-1]

	// loop column by column
	for x := 0; x < len(main)-1; x++ {
		for y := 0; y < len(chart_numbers); y++ {
			letter := main[x]
			future_letter := main[x+1]

			line_index := len(chart_numbers) + lowest_value - y - 1
			symbol := get_symbol(letter, future_letter, line_index)

			var spaces string
			if symbol == "" {
				spaces = " "
			} else {
				spaces = ""
			}

			chart_numbers[y] = fmt.Sprintf("%s%s%s", chart_numbers[y], spaces, symbol)
		}
	}

	// print array
	for _, elem := range chart_numbers {
		fmt.Println(elem)
	}
}

func main() {
	// get currency and sort it
	items := []int{-2, 3, 3, 7, 9, 15, 8, 11, 4, 14, 2}
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
