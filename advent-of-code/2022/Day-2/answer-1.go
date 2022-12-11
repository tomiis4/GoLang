package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
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
	
	// logic
	var score int = 0
	points := map[string]int {
		"A": 1, 
		"B": 2, 
		"C": 3, 
		"win": 6, 
		"draw": 3, 
		"lost": 0,
	}
	objects := map[string]string {
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
	
	
	for _, value := range data {
		items := strings.Split(value, "")
		oponent := items[0]
		me := objects[items[2]]
		
		score += points[me]
		// draw
		if oponent == me {
			score += points["draw"]
		} 
		
		// win
		if (oponent == "A" && me == "B") || (oponent == "B" && me == "C") || (oponent == "C" && me == "A") {
			score += points["win"]
		} 
		
	} 
	fmt.Println(score)
}
