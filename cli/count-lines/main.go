/***********************************************
**                                            **
**   Author: tomiis4                          **
**   Github: https://github.com/tomiis4       **
**   Idea: https://github.com/AlDanial/cloc   **
**                                            **
***********************************************/


package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type LinesCount struct {
	language string
	comments int
	blank    int
	lines    int
}

type File struct {
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

func starts_with_arr(str string, arr []string) bool {
	for _, elem := range arr {
		if strings.HasPrefix(str, elem) {
			return true
		}
	}

	return false
}

func split(str, split string) []string {
	return strings.Split(str, split)
}

func t_print(arr []LinesCount) {
	top_str := "Language | Lines | Blank | Comments"

	fmt.Println(repeat(len(top_str), "-"))
	fmt.Println(top_str)
	fmt.Println(repeat(len(top_str), "-"))

	for _, elem := range arr {
		lang_len := len(elem.language)
		lines_len := len(strconv.Itoa(elem.lines))
		blank_len := len(strconv.Itoa(elem.blank))

		lang_spaces := repeat(11 - lang_len, " ")
		format_lang_lines := fmt.Sprintf("%s%s%d", elem.language, lang_spaces, elem.lines)

		blank_spaces := repeat(19 - (lines_len+lang_len+len(lang_spaces)), " ")
		format_blank := fmt.Sprintf("%s%d", blank_spaces, elem.blank)

		comment_spaces := repeat(27-(blank_len + lines_len + lang_len + len(lang_spaces) + len(blank_spaces)), " ")
		format_command := fmt.Sprintf("%s%d", comment_spaces, elem.comments)

		formated_str := fmt.Sprintf("%s%s%s", format_lang_lines, format_blank, format_command)
		fmt.Println(formated_str)
	}

	fmt.Println(repeat(len(top_str), "-"))
}

func is_file_valid(name string) bool {
	valid := []string{"ts", "js", "jsx", "tsx", "c", "cpp", "cs", "java", "rs", "md", "txt", "go", "v", "sh", "bat", "py", "lua", "sass", "css", "scss", "html", "vim"}

	for _, elem := range valid {
		if strings.ToLower(name) == elem {
			return true
		}
	}

	return false
}

func get_file_info(content string) File {
	var file File

	regex_obj := regexp.MustCompile(`\r?\n`)
	lines := regex_obj.Split(content, -1)

	// get comments and blank
	for _, line := range lines {
		trimmed := strings.Trim(line, " ")

		// blank lines
		if trimmed == "" {
			file.blank++
		}

		// comments
		comments := []string{ "#", "//", "/*", "--", "%", "<!--" }
		if starts_with_arr(trimmed, comments) {
			file.comments++
		}
	}

	file.lines = len(lines)-1

	return file
}

var stats = []LinesCount{}
func get_files(dir string) []LinesCount {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println("Unable to read directory:", dir)
	}

	// get all files and folders
	for _, file := range files {
		file_name := fmt.Sprintf("%s/%s", dir, file.Name())
		file_name_ext := split(file_name, ".")[ len(split(file_name, "."))-1 ]
		ignore := file.Name() != "node_modules" && file.Name() != ".git"

		if file.IsDir() && ignore {
			get_files(file_name)
		} else if ignore && is_file_valid(file_name_ext) {
			// read file
			content, err := ioutil.ReadFile(file_name)
			data := string(content)

			if err != nil {
				fmt.Println("Unable to read file", file.Name())
			}

			file_lines := get_file_info(data)

			does_contain := -1

			// check if it contain
			for index, status := range stats {
				if status.language == file_name_ext {
					does_contain = index
				}
			}

			// append
			if does_contain == -1 {
				stats = append(stats, LinesCount{
					language: file_name_ext,
					blank: file_lines.blank,
					lines: file_lines.lines,
					comments: file_lines.comments,
				})
			} else {
				stats[does_contain].blank += file_lines.blank
				stats[does_contain].lines += file_lines.lines
				stats[does_contain].comments += file_lines.comments
			}
		}
	}

	return stats
}

func main() {
	items := get_files(".")

	t_print(items)
}
