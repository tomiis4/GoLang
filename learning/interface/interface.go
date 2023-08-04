package main

import (
    "fmt"
)

type num int

func (a *num) add(b num) {
    *a += b
}

func main() {
    var x num = 55
    var y num = 5
    x.add(y)

    fmt.Println(x)
}

// type geometry interface {
//     area() float64
// }
//
// type rect struct {
//     width, height float64
// }
//
// func (r rect) area() float64 {
//     return r.width * r.height
// }
//
// func measure(g geometry) {
//     fmt.Println(g)
//     fmt.Println(g.area())
// }
//
// func main() {
//     r := rect{width: 3, height: 4}
//
//     measure(r)
// }
