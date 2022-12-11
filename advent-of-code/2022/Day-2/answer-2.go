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
		"X": 1, 
		"Y": 2, 
		"Z": 3, 
	}
	status := map[string]string {
		"X": "lost",
		"Y": "draw",
		"Z": "win",
		"win": "A", 
		"draw": "B", 
		"lost": "C",
	}
	
	
	for _, value := range data {
		items := strings.Split(value, "")
		oponent := items[0]
		roundType := status[items[2]]
		
		if roundType == "draw" {
			score += 3 

			switch oponent {
				case "A":
					score += points["X"]
				case "B":
					score += points["Y"]
				case "C":
					score += points["Z"]
			}
		}
		
		if roundType == "win" {
			score += 6
			switch oponent {
				case "A":
					score += points["Y"]
				case "B":
					score += points["Z"]
				case "C":
					score += points["X"]
			}
		}	
		
		if roundType == "lost" {
			score += 0
			switch oponent {
				case "A":
					score += points["Z"]
				case "B":
					score += points["X"]
				case "C":
					score += points["Y"]
			}
		}
	} 
	fmt.Println(score)
}
