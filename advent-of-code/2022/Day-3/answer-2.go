package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func getDifferent(arr1 []string, arr2 []string, arr3 []string) []string {
	result := []string{}
	
	for _, value := range arr1 {
		isContain := strings.Contains(strings.Join(arr2, ""), value)
		isContain2 := strings.Contains(strings.Join(arr3, ""), value)
		isPrev := strings.Contains(strings.Join(result, ""), value)
		
		if isContain && isPrev == false && isContain2 {
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
	mainObj := [][]string{}
	
	i := -1
	looped := -1
	for _, value := range data {
		i++
		if (i%3) == 0 {
			looped++
		}
		
		if looped > len(mainObj)-1 {
			val := []string{value}
			mainObj = append(mainObj, val)
		} else {
			mainObj[looped] = append(mainObj[looped], value)
		}
	} 
	
	// logic
	var result int = 0
	alphabet :=  []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z", "A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	
	for _, arr := range mainObj {
		arr1 := strings.Split(arr[0], "")
		arr2 := strings.Split(arr[1], "")
		arr3 := strings.Split(arr[2], "")
		differentArr := getDifferent(arr1, arr2, arr3)
		
		for _, value := range differentArr {
			result += strings.Index(strings.Join(alphabet, ""), value)+1
		}
	} 
	fmt.Println(result)
}
