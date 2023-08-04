# <a href="https://go.dev">Go</a> Cheatsheet <img width="50em" src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/Go_gopher_favicon.svg/2048px-Go_gopher_favicon.svg.png">

<table>
<td>

* [File](#file)
	* [go.mod](#go.mod)
* [Hello world](#hello-world)
* [Importing packages](#importing-packages)
* [Variables & Types](#variables)
	* [structs](#structs)
	* [map](#map)
    * [generics](#generics)
* [Functions](#functions)
* [Logic statements](#logic-statements)
	* [if/else](#ifelse)
	* [switch/case](#switchcase)

</td>
<td>

* [Loop](#loop)
	* [for-i](#for-i)
	* [for-in](#for-in)
	* [while](#while)
* [Converting](#converting)
* [Build-in iunctions](#build-in-functions)
	* [append](#append)
	* [length](#length)
	* [panic](#panic)
	* [copy](#copy)

</td>
<td>

* [Pointers](#pointers)
* [Unit testing](#unit-testing)
* [External file](#external-file)

</td>
<td>

* [Remote packages](#remote-packages)
* [Packages](#packages)
	* [fmt](#fmt)
	* [io/ioutil](#ioioutil)
	* [regexp](#regexp)
    * [net/http](#nethttp)
	* [time](#time)
	* [math/rand](#mathrand)
	* [testing](#testing)
	* [reflect](#reflect)
* [Project ideas](#project-ideas)
</td>
</table>


## File

### Run file
`go run file.go`

### Generate exe file
`go build file.go`

### go.mod
`go mod init <module-url>`


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

// import as <name>
import <name> "package"
```


## Variables
```go
// automatic type
foo := "String"
slice := []<type>{}

// constants
const var <type> = <value>

// set own type
var <name> <type> = <value>
var slice = []<type>{}

// arrays
array := [...]<type>{} // fixed size
array := [][]<type>{} // 2D array

// maps (similar to objects/json in JavaScript)
maps := map[<key-type>]<value-type>{}

/*
Type: 
	bool              = %t = true, false
	int8, 16, 32, 64  = %d = number in range of x bits, can be negative
	uint8, 16, 32, 64 = %d = number in range of x bits, can't be negative
	float32, 64       = %g = decimal numbers
	string            = %s = string
*/
```

### Structs
```go
type Struct struct {
	x <type>
	y <type>
}

// 1
foo := Struct{
	x: <value>,
	y: <value>
}

// 2
foo := Struct{}

foo.x = <value>
foo.y = <value>
```

### Map
```go
// initialize
maps := map[<key-type>]<value-type>{}

// declare
// in this example KEY-TYPE = string, VALUE-TYPE = int
maps := map[<key-type>]<value-type>{
	"bar": 10,
	"foo": 5
}

// add new key
map[<key>] = <value>

// get value
value := map[<key>]

// check if value exists (ok = true|false)
value, ok := map[<key>]

// delete key & value
delete(map, <key>)
```

### Generics
```go
```


## Functions
```go
func name() {
	//...
}

// return
func name() <type> { return x }
func name() (<type>, <type>) { return x, y }

// parameters 
func name(param1 <type>) {  }
func name(param1, param2 <type>) {  } // if param1 have same type as param2

// function for type
func (a <type>) name() {  }
x.name()
```


## Logic Statements

### If/else
```go
if statement {
	//...
} else if statement2 {
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
for i:=0; i<5; i++ {
	//...
}
```

### For-In
```go
for index, value := range arr {
	//...
}
```

### While
```go
for {
	if statement {
		break
	}
}
```


## Converting
```go
import "strconv"

// str -> int
num, err := strconv.Atoi( <string> )

// int -> str
str, err := strconv.Itoa( <number> )

// number -> int8, 16, 32, 64
num := int<bit>( <number> )

// number -> uint8, 16, 32, 64
num := uint<bit>( <number> )

// number -> float32, float64
num := float<bit>( <number> )
```


## Build-In Functions

### Append
```go
<slice> = append(<slice>, <value>)
```

### Length
```go
// slice
slice := []uint8{1,2,3,4}
sliceLength := len(slice) // 4

// string
str := "hello"
strLength := len(str) // 5

// map
maps := map[string]int{ "x":10, "y":15 }
mapsLength := len(maps) // 2
```

### Panic
```go
// make runtime error and stops the program
panic( <message> )
```

### Copy
```go
copy(<slice-to>, <slice-from>)
```


## Pointers
```go
// create new pointer and allocate memory
pointer := new(<type>)

// change data from address/pointer
*pointer = <value>

// read data from address/pointer
*pointer

// get address from variable
&variable
```


## Unit Testing
```sh
go test
```
```go
// main.go
package main

func abs(number int) int {
	if number > 0 { return number }
	return number * -1
}

// main_test.go
// test files must have _test
package main

import "testing"

// starts with capital letter
func TestAbs(t *testing.T) {
	if abs(-1) < 0 {
		t.Error("Negative value was found in abs() with", -1)
	}
}
```


## External file
```go
// create go.mod if you don't have one
go mod init <link-to-module> // for this we will use 'modules'

// folder structure
|- go.mod
|- main.go
|
|- example
  |- second.go

// main.go
package main
import "modules/example"

func main() {
	example.Foo()
}

// second.go
package example

func Foo() {
	//...
}
```



## Remote packages

### Install packages
```
go get <link-to-module>
```

### Import packages
```go
import "<link-to-module>"
```


## Packages

### fmt

#### Print content
```go
import "fmt"

// print on new line, variables using next argument
fmt.Println(...)

// print on same line, variables using format
fmt.Printf(...)
```

#### Get user input
```go
import "fmt"

var variable <type>
fmt.Scanf("%<format>", &variable)
```

#### Format string
```go
import "fmt"

fmt.Sprintf("%<format> %<format>", <variable>, <variable>)
```

### io/ioutil

#### Read file content
```go
import "io/ioutil"

func main() {
	// read file and save it to variable data
	data, err := ioutil.ReadFile("./file")
	
	// check for errors
	if err != nil { panic(err) }
}
```

### regexp

#### Split string using regex
```go
import "regexp"

func main() {
	// regex pattern for end of the line
	pattern := `\r?\n` 
	regex_compiled := regexp.MustCompile(pattern)
	
	// split and save it to variable data
	data := regex_compiled.Split("string", -1) // param1=string, param2=how many time do action 
}
```

### net/http

##### Simple http server
```go
import (
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        // http://127.0.0.1:<port>/bar
    })

    port := "8080"
    err := http.ListenAndServe(":" + port, nil)

    if err != nil {
        log.Fatal(err)
    }
}
```

#### HTTP Requests
```go
import "net/http"

// GET request
resp, err := http.Get(url)
defer resp.Body.Close() // close connection
body, err := io.ReadAll(resp.Body)

// POST request
resp, err := http.Post(url, content_type, &buf)
defer resp.Body.Close() // close connection
body, err := io.ReadAll(resp.Body)
```

### time

#### Delay
```go
import "time"

// ms
func delayMs(ms time.Duration) {
	time.Sleep(ms * time.Milisecond)
}

// sec
func delaySec(s time.Duration) {
	time.Sleep(s * time.Second)
}
```

### math/rand

#### Random number
```go
import (
	"time"
	"math/rand"
)

func randInt(maxNumber int) int {
	// reset time, so it will be random
	newTime := rand.NewSource(time.Now().UnixNano())
	resetRandom := rand.New(newTime)
	
	// get random number
	random_number := resetRandom.Intn(maxNumber)
	
	return random_number
}
```

### testing

#### methods
```go
import "testing"

func TestFunction(t *testing.T) {
	// throw error and stop
	t.Error("message")

	// throw error and contiune
	t.Fail("message")

	// print message
	t.Log("message")
	t.Logf("message %d", 1)
}
```

### reflect

#### type-of
```go
import "reflect"

// get type of variable
reflect.TypeOf(<variable>) // return type
```


## Todo
- [ ] goroutines/channels
- [ ] add types, byte..
- [ ] stdlib -> encoding, sort, strconv, strings, json
- [ ] generics
- [ ] interface


## Project ideas
* [2048 game](https://github.com/tomiis4/GoLang/tree/main/2048)
* [language guessing game](https://github.com/tomiis4/GoLang/tree/main/guess-lang)
* [cli plot chart](https://github.com/tomiis4/GoLang/tree/main/cli/plot-chart)
* [simple cli tower defense](https://github.com/tomiis4/GoLang/tree/main/tower-defense)
* [calculator](https://github.com/tomiis4/GoLang/tree/main/calculator)
* [rock-paper-scissors](https://github.com/tomiis4/GoLang/tree/main/rps)
