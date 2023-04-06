package main

import (
   "fmt"
   "strconv"
   "strings"
)

const (
   LENGTH int   = 21
   WHITE string = " "
   BLACK string = "█"
)

func repeat(s string, n int) string {
   str := ""

   for i:=0; i < n; i++ {
      str += s
   }

   return str
}

func isInt(str string) bool {
    _, err := strconv.Atoi(str)
    return err == nil
}

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

func toAlphanumerics(s string) int {
   if isInt(s) {
      num, _ := strconv.Atoi(s)
      return num
   }

   switch s {
   case "A":
      return 10
   case "B":
      return 11
   case "C":
      return 12
   case "D":
      return 13
   case "E":
      return 14
   case "F":
      return 15
   case "G":
      return 16
   case "H":
      return 17
   case "I":
      return 18
   case "J":
      return 19
   case "K":
      return 20
   case "L":
      return 21
   case "M":
      return 22
   case "N":
      return 23
   case "O":
      return 24
   case "P":
      return 25
   case "Q":
      return 26
   case "R":
      return 27
   case "S":
      return 28
   case "T":
      return 29
   case "U":
      return 30
   case "V":
      return 31
   case "W":
      return 32
   case "X":
      return 33
   case "Y":
      return 34
   case "Z":
      return 35
   case " ":
      return 36
   case "$":
      return 37
   case "%":
      return 38
   case "*":
      return 39
   case "+":
      return 40
   case "-":
      return 41
   case ".":
      return 42
   case "/":
      return 43
   case ":":
      return 44
   default:
      return 0
   }
}

func strToBin(str string) string {
   bin := ""

   // group for 2 chars
   arr := strings.Split(str, "")
   for i:=0; i < len(str)-1; i+=2 {
      char1 := toAlphanumerics(arr[i])
      char2 := toAlphanumerics(arr[i+1])

      binaryStr := strconv.FormatInt(int64(char1*45 + char2), 2)
      paddLen := 11 - len(binaryStr)

      binaryStr = repeat("0", paddLen) + binaryStr

      bin += binaryStr
   }

   // if is one untouched
   if len(str) % 2 != 0 {
      char := toAlphanumerics(arr[len(str)-1])
      binaryChar := strconv.FormatInt(int64(char), 2)

      paddLen := 6 - len(binaryChar)

      bin += repeat("0", paddLen) + binaryChar
   }

   return bin
}

func getTerminator(size, max int) string {
   if size+4 <= max {
      return "0000"
   } else if size+3 == max {
      return "000"
   } else if size+2 == max {
      return "00"
   } else if size+1 == max {
      return "0"
   } else {
      return ""
   }
}

func fillBytes(current string, max int) string {
   filled := current

   // fill to even numbers
   if len(current) % 8 != 0 {
      filled += repeat("0", 8 - len(current)%8)
   }

   // fill padd bytes
   paddBytes := [...]string{"11101100", "00010001"}
   paddLen := (max - len(filled)) / 8

   for i:=0; i < paddLen; i++ {
      filled += paddBytes[i % 2]
   }

   return filled
}

func splitDecimals(code string) []int {
   decimals := []int{}

   for i := 0; i < len(code); i += 8 {
      end := i + 8
      if end > len(code) {
         end = len(code)
      }

      decimal, _ := strconv.ParseInt(code[i:end], 2, 64)
      decimals = append(decimals, int(decimal))
   }

   return decimals
}

// alphanumeric mode
// indicator = 0010
// https://www.thonky.com/qr-code-tutorial/error-correction-table
func encodeStr() {
   maxBytes := 16 * 8// (1-M)
   str := "HELLO WORLD"
   strLen := len(str)

   strLenBinary := strconv.FormatInt(int64(strLen), 2)

   modeIndicator := "0010"
   charCountIndicator := fmt.Sprintf("%09s", strLenBinary)
   encodedData := strToBin(str)
   terminator := getTerminator(4+len(charCountIndicator)+len(encodedData), maxBytes)

   filled := fillBytes(modeIndicator+charCountIndicator+encodedData+terminator, maxBytes)
   
   decimals := splitDecimals(filled)

   fmt.Println(decimals)
}

func main() {
   //code := generateBlank()
   //addDefaults(code)

   //printCode(code)
   encodeStr()
}
