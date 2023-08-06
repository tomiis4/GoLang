package main

import (
	"fmt"
	"reflect"
)

func main() {
	// cap 
	//		cap == len ?
	arr := [10]int{}
	arr[0] = 1
	arr[5] = 10

	fmt.Println(arr, cap(arr), len(arr))

	// copy
	//		value is copied to first argument,
	//		if second argument is smaller, 
	//     argument 1 still have old value
	//		 example: copy({1,2,3}, {5,6}) -> {5,6,3}
	val1 := []int{1,2,3}
	val2 := []int{5,6}

	copy(val1, val2)

	fmt.Println(val1, val2)

	// append
	//		return slice where you appended value
	//		slice1 = {8,9,7,5}
	slice1 := []int{1,2,3}
	slice2 := []int{8,9,7}

	slice1 = append(slice2, 5)

	fmt.Println(slice1, slice2)

	// map
	//		as array have index, map have key in loop
	//		if you are getting element by key, you are getting 2 variables, 1=value, 2=if it exist
	maps := map[string]int{ "x":10, "y":15 }

	for key, value := range maps {
		fmt.Println(key, value)
	}
	val, ok := maps["x"]

	fmt.Println(val, ok)

	// reflect
	//		check if types work
	var str string = "hello"
	var num64 int64 = 5645412 

	fmt.Println(str, ":", reflect.TypeOf(str), "\n", num64, ":", reflect.TypeOf(num64))

    // convert
    bytes := []byte("Ahoj")
    fmt.Println(string(bytes), bytes)
}
