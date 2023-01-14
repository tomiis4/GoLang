package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
	file "github.com/tomiis4/getFile"
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

func getIndex(arr []string, elem string) int {
	for i, value := range arr {
		if value == elem {
			return i
		}
	}
	return -1
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
	languages := [18]string{"Java", "C", "Python", "TypeScript", "C++", "C#", "P*P", "JavaScript", "SQL", "Ruby", "Matlab", "Swift", "Go", "COBOL", "Fortran", "Rust", "Haskell", "Bash"}
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
	fmt.Println("\n",strSpace, "\n ")
	// code
	displayCode(code, codeRevaled)
	fmt.Println("\n",strSpace, "\n ")
	// languages
	displayLanguages(languages)
	fmt.Println(strSpace)
}

func clearConsole() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// lang
	currentLang := "TypeScript"
	// currentLang := "C"
	languages := getLanguages(currentLang)
	langIndex := getIndex(languages, currentLang) 
	//code
	code := file.GetContent("typescript")
	// code := []string{"#include <stdio.h>", "int main(", "  printf(Helloworld);", "  return 0;"} 
	codeRevaled := uint8(0)
	// score
	score := uint16(0)
	questionScore := uint8(250)
	
	// game loop
	inputChannel := make(chan int)
	
	// goroutine input
	go func() {
		// scan for input
		for ;; {
			var numberInput int 
			fmt.Scanf("%d", &numberInput)
			// send data to another thread
			inputChannel <- numberInput
		}
	}()
	
	// goroutine game-loop 
	go func() {	
		for ;; {
			clearConsole()
			displayUI(score, questionScore, code, codeRevaled, languages)
			
			// change
			codeRevaled++
			if questionScore-25 >0 { questionScore-=25 }
			delay(1500)
			
			// check for input
			select {
				case numberInput := <- inputChannel:
					// if you guess corrent lang
					if numberInput == langIndex+1 {
						// add score
						score += uint16(questionScore)
						questionScore = 250
						codeRevaled = 0
						code = file.GetContent("typescript")
						languages = getLanguages(currentLang)
						langIndex = getIndex(languages, currentLang) 
					}
				default:
					// nothing (this should never activate, because int will be always 0 I guess)
			}
		}
	}()
	
	// check for end
	select {}
}
