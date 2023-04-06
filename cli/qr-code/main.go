package main

import (
   "fmt"
   "strconv"
)

const (
   LENGTH int   = 21
   WHITE string = " "
   BLACK string = "█"
)

func Pow(a,b int) int {
   res := 1

   for i:=0; i < b; i++ {
      res *= a
   }

   return res
}

func getIndex(x,y int) int {
   return LENGTH * (y) + (x)
}

func addDefaults(code []string) {
   square := [...][]int{
      {0, 0}, {1,0}, {2,0}, {3,0}, {4,0}, {5,0}, {6,0},
      {0, 6}, {1,6}, {2,6}, {3,6}, {4,6}, {5,6}, {6,6},
      {0, 1}, {6,1}, {0, 2}, {6,2}, {0, 3}, {6,3}, {0, 4}, {6, 4}, {0, 5}, {6,5}, {0, 6}, {6,6},
      {2,2}, {3,2}, {4,2}, {2,3}, {3,3}, {4,3}, {2,4}, {3,4}, {4,4},
   }

   // for 3x squares
   for i:=0; i < 3; i++ {
      addX := 0
      addY := 0

      // add 2 squares
      if i == 1 {
         addX = 14
         addY = 0
      } else if i == 2 {
         addX = 0
         addY = 14
      }

      for _, vx := range square {
         index := getIndex(vx[0]+addX, vx[1]+addY)

         code[index] = WHITE
      }

   }

   // add default items
   defaultSqr := [...][]int{
      {8,6}, {10,6}, {12,6},
      {6,8}, {6,10}, {6,12},
   }

   for _, vx := range defaultSqr {
      index := getIndex(vx[0], vx[1])

      code[index] = WHITE
   }

}

func generateBlank() []string {
   blank := []string{}

   for i:=0; i < Pow(LENGTH, 2); i++ {
      blank = append(blank, BLACK)
   }

   return blank
}

func printCode(code []string) {
   for index, value := range code {
      fmt.Printf("%s", value)

      if index % LENGTH == LENGTH-1 {
         fmt.Printf("\n")
      }
   }
}

func strToBin(str string) string {
   bin := ""

   // group for 2 chars
   for i:=0; i < len(str)-1; i+=2 {
      char1 := int(str[i]) * 45
      char2 := int(str[i+1])

      binaryStr := strconv.FormatInt(int64(char1 + char2), 2)
      paddLen := 11 - len(binaryStr)

      for padd:=0; padd < paddLen; padd++ {
         binaryStr = "0" + binaryStr
      }

      bin += binaryStr + " "
   }

   return bin
}

// alphanumeric mode
// indicator = 0010
func encodeStr() {
   str := "HELLO WORLD"
   strLen := len(str)

   strLenBinary := strconv.FormatInt(int64(strLen), 2)
   paddedBinary := fmt.Sprintf("%09s", strLenBinary)

   binCount := "0010" + paddedBinary
   binStr := strToBin(str)

   fmt.Println(binStr)
   fmt.Println(binCount)
}

func main() {
   //code := generateBlank()
   //addDefaults(code)

   //printCode(code)
   encodeStr()
}
