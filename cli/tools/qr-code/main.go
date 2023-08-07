// ******************************* //
// *                             * //
// *         Informations        * //
// *                             * //
// * thonky.com/qr-code-tutorial * //
// *                             * //
// ******************************* //


package main

import (
   "fmt"
   "strconv"
   "strings"
   enc "tomiis4/qr/encoder"
)

const (
   LENGTH   int   = 21
   WHITE    string = " " // 0
   BLACK    string = "█" // 1
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
   finderPattern := [...][]int{
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

      for _, vx := range finderPattern {
         index := getIndex(vx[0]+addX, vx[1]+addY)

         code[index] = WHITE
      }

   }

   // add timing patterns
   timingPatterns := [...][]int{
      {8,6}, {10,6}, {12,6},
      {6,8}, {6,10}, {6,12},
   }

   for _, vx := range timingPatterns {
      index := getIndex(vx[0], vx[1])

      code[index] = WHITE
   }

   // add dark module
   darkModule := [...]int{8, (4 * 1) + 9}
   darkModuleIndex :=getIndex(darkModule[0], darkModule[1])

   code[darkModuleIndex] = WHITE
}

func generateBlank() []string {
   blank := []string{}

   for i:=0; i < Pow(LENGTH, 2); i++ {
      blank = append(blank, BLACK)
   }

   return blank
}

func printCode(code []string) {
   fmt.Println(repeat("-", LENGTH))
   for index, value := range code {
      fmt.Printf("%s", value)

      if index % LENGTH == LENGTH-1 {
         fmt.Printf("\n")
      }
   }
   fmt.Println(repeat("-", LENGTH))
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

func intsToBin(dec []int) string {
   binary := ""

   for _, num := range dec {
      bin := strconv.FormatInt(int64(num), 2)
      binary += fmt.Sprintf("%08s", bin)
   }

   return binary
}

// alphanumeric mode
// indicator = 0010
// https://www.thonky.com/qr-code-tutorial/error-correction-table
func encodeStr() string {
   maxBytes := 16 * 8 // (1-M)
   //codewords := 10 // (byte)
   str := "HELLO WORLD"
   strLen := len(str)

   strLenBinary := strconv.FormatInt(int64(strLen), 2)

   modeIndicator := "0010"
   charCountIndicator := fmt.Sprintf("%09s", strLenBinary)
   encodedData := strToBin(str)
   terminator := getTerminator(4+len(charCountIndicator)+len(encodedData), maxBytes)

   filled := fillBytes(modeIndicator+charCountIndicator+encodedData+terminator, maxBytes)

   decimals := splitDecimals(filled)
   codewordsDecimal := enc.Encoder(decimals)
   codewords := intsToBin(codewordsDecimal)

   data := filled + codewords


   return data
}

func addData(code []string, data string) {
   dataArr := strings.Split(data, "")
   indexOrder := []int{440, 439, 419, 418, 398, 397, 377, 376, 356, 355, 335, 334, 314, 313, 293, 292, 272, 271, 251, 250, 230, 229, 209, 208, 207, 206, 228, 227, 249, 248, 270, 269, 291, 290, 312, 311, 333, 332, 354, 353, 375, 374, 396, 395, 417, 416, 438, 437, 436, 435, 415, 414, 394, 393, 373, 272, 352, 351, 331, 330, 310, 309, 289, 288, 268, 267, 247, 246, 226, 225, 205, 204, 203, 202, 224, 223, 245, 244, 266, 265, 287, 286, 308, 307, 329, 328, 350, 349, 371, 370, 392, 391, 413, 412, 434, 433, 432, 431, 411, 410, 390, 389, 369, 368, 348, 347, 327, 326, 306, 305, 285, 284, 264, 263, 243, 242, 222, 221, 201, 200, 180, 179, 159, 158, 117, 116, 96, 95, 75, 74, 54, 53, 33, 32, 12, 11, 10, 9, 31, 30, 52, 51, 73, 72, 94, 93, 115, 114, 157, 156, 178, 177, 199, 198, 220, 219, 241, 240, 262, 261, 283, 282, 304, 303, 325, 324, 346, 345, 367, 366, 388, 387, 409, 408, 430, 429, 260, 259, 239, 238, 218, 217, 197, 196, 194, 193, 215, 214, 236, 235, 257, 256, 255, 254, 234, 233, 213, 212, 192, 191, 190, 189, 211, 210, 232, 231, 253, 252}

   for i, order := range indexOrder {
      bit := dataArr[i]

      if bit == "0" {
         code[order] = WHITE
      } else if bit == "1" {
         code[order] = BLACK
      }
   }
}

func main() {
   code := generateBlank()
   data := encodeStr()

   addDefaults(code)
   addData(code, data)

   printCode(code)
}
