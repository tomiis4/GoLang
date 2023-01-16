# Go Cheatsheet

<table>
<td>

* [Hello world](#hello-world)
* [Importing packages](#importing-packages)
* [Variables & Types](#variables)
	* [objects](#objects)
* [Functions](#functions)
* [Check](#check)
	* [if/else](#ifelse)
	* [switch/case](#switchcase)
</td>
<td>

* [Loop](#loop)
	* [for-i](#for-i)
	* [for-in](#for-in)
	* [while true](#while-true)
* [Build-in Functions](#build-in-functions)
	* [convert number](#convert-number-type)
	* [length](#length)

</td>
<td>

* [Packages](#packages)
	* [io/ioutil](#ioioutil)
	* [regexp](#regexp)
	* [time](#time)
	* [math/rand](#mathrand)
* [Project ideas](#project-ideas)
</td>
</table>

## Hello world
```go
// name file package
package main

// import input/output package
import "fmt"

// main function where will all code execute
func main() {
	fmt.Println("Hello, world")
}
```

## Importing packages
```go
// import one package
import "package"

// import multiple packages
import (
	"package1"
	"package2"
)

// import as X
import pcg "package"
```

## Variables
```go
// DYNAMIC

// automatic type
foo := "String"
slice := []<type>{}

// own type
var <name> <type> = <value>
var slice = []<type>{}

// FIXED
const var <type> = <value>


/*
Type: 
	bool               = true, false
	int8, 16, 32, 64   = number in range of x bits, can be negative
	uint8, 16, 32, 64  = number in range of x bits, can't be negative
	float32, 64        = decimal numbers
*/
```

### Objects
```go
type Object struct {
	x <type>
	y <type>
}

func function() {
	var obj Object
	
	obj.x = <value>
}
```

## Functions
```go
func name() {
	//...
}

// return & argument
func name(arg <type>) <type> {
	//...
	return arg
}
```

## Check
### If/else
```go
if statement {
	//...
} else if statement2{
	//...
} else {
	//...
}
```

### Switch/case
```go
switch statement {
	case x:
		//...
	case y:
		//...
	default:
		//...
}
```

## Loop
### For-I
```go
for i:=0; i < 5; i++ {
	//...
}
```

### For-In
```go
for index, value := range arr {
	//...
}
```

### While true
```go
for ;; {
	if statement {
		break
	}
}
```

## Build-In Functions
### Convert number type
```go
x := uint8(2)
y := float32(3.1415926535)
```

### Length
```go
// array
arr := []uint8{1,2,3,4}
arrayLength := len(arr) // 4

// string
str := "hello"
strLength := len(str) // 5
```

## Packages
### io/ioutil
#### Read file content
```go
import "io/ioutil"

func main() {
	// read file and save it to variable data
	data, err := ioutil.ReadFile("./file")
	
	// check for errors
	if err != nil { fmt.Println(err) }
}
```

### regexp
#### Split string using regex
```go
import "regexp"

func main() {
	// regex pattern for end of the line
	pattern := `\r?\n` 
	regexCompiled := regexp.MustCompile(pattern)
	
	// split and save it to variable data
	data := regexCompiled.Split("string", -1) // arg1 = string, arg2 = how many time do action 
}
```

### time
#### Delay
```go
import "time"

// ms
func delay(ms time.Duration) {
	time.Sleep(ms * time.Milisecond)
}

// second
func delay(s time.Duration) {
	time.Sleep(s * time.Second)
}
```

### math/rand
#### Random number
```go
import "time"
import "math/rand"

func randomInt(maxNumber int) int {
	// give random new time because it will not be always random
	newTime:= rand.NewSource(time.Dow().UnixNano())
	resetRandom := rand.New(newTime)
	
	// get random number
	randomNumber := resetRandom.Intn(maxNumber)
	
	return randomNumber
}
```

## Project ideas
* [2048 game](https://github.com/tomiis4/GoLang/tree/main/2048)
* [language guessing game](https://github.com/tomiis4/GoLang/tree/main/guess-lang)
* [calculator](#)
