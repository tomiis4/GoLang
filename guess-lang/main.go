package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func replaceStr(str string, with string) string {
	strArr := strings.Split(str, "")
	newArr := []string{}
	
	for _, elem := range strArr {
		if elem != " " {
			newArr = append(newArr, with)
		} else {
			newArr = append(newArr, elem)
		}
	}
	
	return strings.Join(newArr, "")
}

func isContain(arr []string, checkElem string) bool {
	for _, elem := range arr {
		if elem == checkElem {
			return true
		}
	}
	return false
}

func randInt(max int) int {
	randTime := rand.NewSource(time.Now().UnixNano())
	newRand := rand.New(randTime)
	randomNum := newRand.Intn(max)
	
	return randomNum
}

func delay(ms time.Duration) {
	time.Sleep(ms * time.Millisecond)
}

func getLine(strArr []string) string {
	//FIXME need fix tho
	letterLen := len(strings.Split(strings.Join(strArr, " "), "")) *2+3
	line := []string{}
	
	
	for i:=0; i < letterLen; i++ {
		line = append(line, "-")
	}
	
	return strings.Join(line, "")
}

func displayCode(codeBlock []string, level uint8) {
	for i:=uint8(0); i < uint8(len(codeBlock)); i++ {
		elem := codeBlock[i]
		
		// display visible lines
		if i < level {
			fmt.Println(i, elem)
		} else {
			// display hidden lines
			fmt.Println(i, replaceStr(elem, "*"))
		}
	}
}

func getLanguages(current string) []string {
	languagesChoices := 4
	languages := [17]string{"Java", "C", "Python", "C++", "C#", "P*P", "JavaScript", "SQL", "Ruby", "Matlab", "Swift", "Go", "COBOL", "Fortran", "Rust", "Haskell", "Bash"}
	selectedLanguages := []string{}
	
	// create arr of langs
	for ;; {
		// random lang
		randomNum := randInt(len(languages))
		lang := languages[randomNum]
		
		if !isContain(selectedLanguages, lang) && len(selectedLanguages) < languagesChoices && lang != current {
			selectedLanguages = append(selectedLanguages, lang) 
		}
		
		if len(selectedLanguages) == languagesChoices {break}
	}
	
	// insert correct language
	selectedLanguages[randInt(languagesChoices)] = current
	
	return selectedLanguages
}

func displayLanguages(languages []string) {
	for i, elem := range languages {
		if i == 0 { fmt.Printf("|") }
		fmt.Printf(" %d. %s |", i+1, elem)
		
		if i+1 == len(languages) { fmt.Printf("\n") }
	}
}

func displayUI(score uint16, questionScore uint8, code []string, codeRevaled uint8, languages []string) {
	strSpace := getLine(languages)	
	
	fmt.Println(strSpace)
	// score
	fmt.Println("Total score:", score)
	fmt.Println("Current score:", questionScore)
	fmt.Println(strSpace)
	// code
	displayCode(code, codeRevaled)
	fmt.Println(strSpace)
	// languages
	displayLanguages(languages)
	fmt.Println(strSpace)
}

func main() {
	code := []string{"#include <stdio.h>", "int main(", "  printf(Helloworld);", "  return 0;"} 
	codeRevaled := uint8(0)
	currentLang := "C"
	languages := getLanguages(currentLang)
	score := uint16(2000)
	questionScore := uint8(250)
	
	
	for ;; {
		displayUI(score, questionScore, code, codeRevaled, languages)
		codeRevaled++
		delay(1500)
	}
	
}
