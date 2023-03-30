package main

import "fmt"


func sort(data []int) {
	isSorted := false

	for i:=0; i<len(data); i++ {
		for j:=0; j<len(data)-1; j++ {
			if data[j] > data[j+1] {
				isSorted = false
				data[j], data[j+1] = data[j+1], data[j]
			}
		}

		if isSorted {
			return
		}
	}
}

func main() {
	data := []int{5,1,3,7,6,2}

	sort(data)

	fmt.Println(data)
}
