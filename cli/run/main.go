package main

import (
	"os"
	"os/exec"
	"strings"
	"fmt"
)

func main() {
   args := os.Args
   extension := strings.Split(args[1], ".")[len(strings.Split(args[1], "."))-1]
   params := strings.Join(args[2:], " ")

   fileName := args[1] + " "

   cmd := exec.Command("echo [empty]")

   switch extension {
   case "exe":
      cmd = exec.Command(fileName + params)
   case "go":
      cmd = exec.Command("go run " + fileName + params)
   case "c":
      cmd = exec.Command("gcc " + fileName + params + "&& a.exe")
   case "js":
      cmd = exec.Command("node " + fileName + params)
   case "ts":
      cmd = exec.Command("ts-run " + fileName + params)
   case "v":
      cmd = exec.Command("v run " + fileName + params)
   case "zig":
      cmd = exec.Command("zig run " + fileName + params)
   default:
      panic(fmt.Sprintf("Format *.%s is not supported.", extension))
   }

   cmd.Run()
}
