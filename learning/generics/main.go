package generics

import "fmt"

func printAny[S interface{}](s []S) {
    for k, v := range s {
        fmt.Println(k, v)
    }
}

func main() {
    printAny([]int{4})
}
