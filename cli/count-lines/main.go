package main

import (
	"fmt"
	"strings"
)

type LinesCount struct {
	language string
	comments int
	blank    int
	lines    int
}

func repeat(n int, item string) string {
	str := []string{}

	for i:=0; i < n; i++ {
		str = append(str, item)
	}

	return strings.Join(str, "")
}

func t_print(arr []LinesCount) {
	fmt.Println(repeat(10, "-"))

	for _, value := range arr {
		fmt.Println(value)
	}

	fmt.Println(repeat(10, "-"))
}

func main() {
	fmt.Println("Hello")
}
