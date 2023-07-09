package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type File struct {
	name     string
	isFolder bool
}

var icons = map[string]string{
	"folder":       "",
	"txt":          "",
	"img":          "",
	"png":          "",
	"jpg":          "",
	"jpeg":         "",
	"audio":        "",
	"video":        "",
	"exe":          "",
	"toml":         "",
	"pdf":          "",
	"md":           "",
	"json":         "",
	"xml":          "謹",
	"html":         "",
	"css":          "",
	"scss":         "",
	"js":           "",
	"ts":           "",
	"LICENSE":      "󰿃",
	"py":           "",
	"rust":         "",
	"java":         "",
	"c":            "",
	"cpp":          "",
	"go":           "",
	"rb":           "",
	"php":          "",
	"sh":           "",
	"bat":          "",
	"docker":       "",
	"sql":          "",
	"vim":          "",
	"nvim":         "",
	"neovim":       "",
	"vimrc":        "",
	"git":          "",
	"gitignore":    "",
	"make":         "",
	"config":       "",
	"vscode":       "",
	"node_modules": "",
}

func getFiles() []File {
	files := []File{}

	// read dir
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	// get files
	for _, v := range entries {
		files = append(files, File{
			name:     v.Name(),
			isFolder: v.IsDir(),
		})
	}

	return files
}

func getFlag() string {
	flag := ""
	arg := os.Args
	if len(arg) > 1 {
		flag = arg[1]
	}

	return strings.Replace(flag, "-", "", -1)
}

func getIcon(name string, isFolder bool) string {
	icon, ok := icons[name]

	if ok {
		return icon
	} else {
		if isFolder {
			return icons["folder"]
		} else {
			return icons["txt"]
		}
	}
}

func displayFiles(files []File, arg string) {
	for i, v := range files {
		name := v.name

		// if display all, not ignore .
		if arg != "a" && string(name[0]) == "." {
			name = ""
			v.isFolder = false
		}

		// color folders
		if v.isFolder {
			icon := getIcon(name, true)

			fmt.Printf("\033[36m"+icon+" %s\033[0m  ", name)
		} else if name != "" {
			ext := strings.Split(name, ".")[len(strings.Split(name, "."))-1]
			icon := getIcon(ext, false)

			fmt.Printf(icon+" %s  ", name)
		}

		if i%3 == 0 && i != 0 {
			fmt.Println()
		}
	}
}

func main() {
	flag := getFlag()
	files := getFiles()

	displayFiles(files, flag)
}
