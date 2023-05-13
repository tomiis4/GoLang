package main

import (
	"os"
	"os/exec"
	"strings"
	"fmt"
)

func main() {
   args := os.Args
   extension := strings.Split(args[0], ".")[len(strings.Split(args[0], "."))-1]
   params := strings.Join(args[1:], "")

   cmd := exec.Command("echo [empty]")

   switch extension {
   case "exe":
      cmd = exec.Command(args[0] + params)
   case "go":
      cmd = exec.Command("go run " + args[0] + params)
   case "c":
      cmd = exec.Command("gcc " + args[0] + params + "&& a.exe")
   case "js":
      cmd = exec.Command("node " + args[0] + params)
   case "ts":
      cmd = exec.Command("ts-run " + args[0] + params)
   case "v":
      cmd = exec.Command("v run " + args[0] + params)
   case "zig":
      cmd = exec.Command("zig run " + args[0] + params)
   default:
      panic(fmt.Sprintf("Format %s is not supported.", extension))
   }

   cmd.Run()
}
