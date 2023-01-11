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

func displayLanguages(current string) {
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
	
	fmt.Println(selectedLanguages)
}

func displayUI(score uint16, thisScore uint8, codeBlock []string, codeLang string) {
	const strSpace string = "------------------------------------"
	
	fmt.Println(strSpace)
	
	fmt.Println("Total score:", score)
	fmt.Println("Current score:", thisScore)
	
	fmt.Println(strSpace)
	
	displayCode(codeBlock, 2)
	
	fmt.Println(strSpace)
	
	displayLanguages(codeLang)
	
	fmt.Println(strSpace)
}

func main() {
	fmt.Println("Hello, world")
	
	code := []string{"#include <stdio.h>", "int main(", "  printf(Helloworld);", "  return 0;"} 
	
	displayUI(1800, 225, code, "C")
}
