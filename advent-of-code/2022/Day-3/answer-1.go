package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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
	alphabet :=  []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z", "A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
	
	// logic
	for _, value := range data {
		items := strings.Split(value, "")
		firstHalf := items[0:(len(items)/2)]
		secondHalf := items[(len(items)/2):]
		
		differentArr := getDifferent(firstHalf, secondHalf)
		
		for _, value := range differentArr {
			result += strings.Index(strings.Join(alphabet, ""), value)+1
		}
	} 
	fmt.Println(result)
}
