package main

import "fmt"

func sort(data []int)  {
	// start from second number
	for i:=1; i<len(data); i++ {
		currentNum := data[i]
		// count from that number to first
		j := i-1
		for j>=0 && data[j] > currentNum {
			data[j+1] = data[j]
			fmt.Println(data)
			j--
		}

		data[j+1] = currentNum
	}
}

func main() {
	data := []int{5,1,3,7,6,2} //1 2 3 5 6 7

	sort(data)

	fmt.Println(data)
}
