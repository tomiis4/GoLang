package main

import (
	"fmt"
	"sort"
)

func get_lowest_value(arr []int) int {
	sort.Ints(arr)

	return arr[0]
}

func get_higest_value(arr []int) int {
	sort.Ints(arr)

	return arr[len(arr)-1]
}


/*

How this will work?
I don't know too..



*/
func cprint(sorted []int, main []int) {
	higestValue := sorted[len(sorted)-1]
	lowestValue := sorted[0]

	for value, index := range sorted {
	}
}

func main() {
	fmt.Println("Get stock0 price for BTC")

	// get currency and sort it
	currency := []int { 24000,25200,12514,13654,12311 }
	sortedCurrency := currency
	sort.Ints(sortedCurrency)

	cprint(sortedCurrency, currency)


	// for i:=0; i < len(currency); i++ {
	// 	fmt.Println(currency[i])
	// }
}
