package main

import (
   "fmt"
)

func generateBlank(size int) []string {
   blank := []string{}

   for i:=0; i<size*size; i++ {
      blank = append(blank, "X")
   }

   return blank
}

func main() {
   size := 21
   fmt.Println(generateBlank(size))
}
