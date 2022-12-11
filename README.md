# Go Cheatsheet

## Helo world
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

## Import
```go
// input/output
import "fmt"
```

## Variables 
```go
var boolean bool = true;
var number int = -5 // int8, int16, int32, int64
var number uint = 5 // uint8, uint16, uint32, uint64
var float float32 = 3.14 // float32, float64
var stringName string = "Hi"

variable := "any type"
const pi float32 = 3.14

array := [2]int{1,2}
var array = [...]int{1,2,3,4}
```

## Function
```go
// input/output
func main(arg1: string) int {
	return 4
}
```

## If/else
```go
if something {
	// do something
} else if something2 {
	// do something2
} else {
	// else
}
```

## Switch/case
```go
switch x {
	case 1,2: // if 1 or 2
		// something
	case 3:
		// something
	default:
		// else
}
```

## Loops
## For i
```go
for i:=0; i<5; i++ {
	// someting
}
```
## For in
```go
for index, value := range array {}
```

## Type 
```go
type Person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var person  Person
	
	person.name = "name"
	//...
}
```

## Object 
```go
var a = map[string]string{"key1": "value1"}
```

## Read file 
```go
import ("fmt" "io/ioutil")

func main() {
	content, err := ioutil.ReadFile("file.txt")
	
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(content)
}
```
