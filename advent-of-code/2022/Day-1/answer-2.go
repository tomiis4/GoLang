package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"sort"
)

func main() {
	// parse data
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var dataStr string = string(content)
	var pattern string = `\r?\n`
	regexObj := regexp.MustCompile(pattern)
	data := regexObj.Split(dataStr, -1)
	
	elves := []int{}	
	
	i := 0
	for _, value := range data {
		// each elv
		if value == "" {
			i++
		} 
		
		// convert str to int and add it to array
		valueNum,_ := strconv.Atoi(value)
		if i > len(elves)-1 {
			elves = append(elves, valueNum)
		} else {
			elves[i] += valueNum
		}
	}
	
	// reverse and sort
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	
	// top 3
	var topThree int = elves[0] + elves[1] + elves[2]
	
	fmt.Println(topThree)
}
