package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func getDifferent(arr1 []string, arr2 []string) []string {
	result := []string{}
	
	for _, value := range arr1 {
		str := strings.Join(arr2, "")
		isContain := strings.Contains(str, value)
		isPrev := strings.Contains(strings.Join(result, ""), value)
		
		if isContain && isPrev == false {
			result = append(result, value)
		}
	}
	
	return result
}

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
	
	var result int = 0
	
	// logic
	for _, value := range data {
		items := strings.Split(value, ",")
		value1 := strings.Split(items[0], "-")
		value2 := strings.Split(items[1], "-")
		
		aVal1,_ := strconv.Atoi(value1[0])
		aVal2,_ := strconv.Atoi(value1[1])
		
		bVal1,_ := strconv.Atoi(value2[0])
		bVal2,_ := strconv.Atoi(value2[1])
		
		if aVal1 <= bVal1 && aVal2 >= bVal2 {
			result++
		}
		if bVal1 <= aVal1 && bVal2 >= aVal2 {
			result++
		} 	
	} 
	fmt.Println(result)
}
