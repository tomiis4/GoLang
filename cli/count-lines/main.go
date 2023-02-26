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
	"time"
)

type LinesCount struct {
	language string
	files    int
	blank    int
	lines    int
}

type File struct {
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

func sum(elemType string, stats []LinesCount) int {
	n := 0

	if elemType == "lines" {
		for _, e := range stats {
			n += e.lines
		}
	}

	if elemType == "blank" {
		for _, e := range stats {
			n += e.blank
		}
	}

	if elemType == "files" {
		for _, e := range stats {
			n += e.files
		}
	}

	return n
}

func t_print(arr []LinesCount, stats []LinesCount) {
	top_str := "Language | Lines | Blank | Files"

	fmt.Println(repeat(len(top_str), "-"))
	fmt.Println(top_str)
	fmt.Println(repeat(len(top_str), "-"))

	// print stats
	for _, elem := range arr {
		lang_len := len(elem.language)
		lines_len := len(strconv.Itoa(elem.lines))
		blank_len := len(strconv.Itoa(elem.blank))

		lang_spaces := repeat(11 - lang_len, " ")
		format_lang_lines := fmt.Sprintf("%s%s%d", elem.language, lang_spaces, elem.lines)

		blank_spaces := repeat(19 - (lines_len+lang_len+len(lang_spaces)), " ")
		format_blank := fmt.Sprintf("%s%d", blank_spaces, elem.blank)

		file_spaces := repeat(27-(blank_len + lines_len + lang_len + len(lang_spaces) + len(blank_spaces)), " ")
		format_file := fmt.Sprintf("%s%d", file_spaces, elem.files)

		formated_str := fmt.Sprintf("%s%s%s", format_lang_lines, format_blank, format_file)
		fmt.Println(formated_str)
	}

	fmt.Println(repeat(len(top_str), "-"))

	// print sum
	sum_lines := sum("lines", stats)
	sum_blank := sum("blank", stats)
	sum_files := sum("files", stats)

	sum_lines_len := len(strconv.Itoa(sum_lines))
	sum_blank_len:= len(strconv.Itoa(sum_blank))

	sum_lines_f := fmt.Sprintf("%s%d", repeat(7, " "), sum_lines)
	sum_blank_f := fmt.Sprintf("%s%d", repeat(19-(11+sum_lines_len)," "), sum_blank)
	sum_files_f := fmt.Sprintf("%s%d", repeat(27-(4+7+sum_lines_len+ (19-(11+sum_lines_len)) +sum_blank_len), " "),sum_files)

	fmt.Printf("sum:%s%s%s\n", sum_lines_f, sum_blank_f, sum_files_f)

	fmt.Println(repeat(len(top_str), "-"))
}

func is_file_valid(name string) bool {
	valid := []string{"ts", "js", "jsx", "tsx", "c", "cpp", "cs", "java", "rs", "md", "txt", "go", "v", "sh", "bat", "py", "lua", "sass", "css", "scss", "html", "vim", "json", "gitignore", "mod", "class", "pyw", "env"}

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

		if trimmed == "" {
			file.blank++
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

		// if is "file" folder then loop trough it
		if file.IsDir() && ignore {
			get_files(file_name)
		} 

		// if it's file and have valid file extension
		if ignore && !file.IsDir() && is_file_valid(file_name_ext) {
			// read file
			content, err := ioutil.ReadFile(file_name)
			data := string(content)

			if err != nil {
				fmt.Println("Unable to read file", file.Name())
			}

			file_lines := get_file_info(data)
			does_contain := -1

			// check if it contain or not
			for index, status := range stats {
				if status.language == file_name_ext {
					does_contain = index
				}
			}

			// append to stats
			if does_contain == -1 {
				stats = append(stats, LinesCount{
					language: file_name_ext,
					blank: file_lines.blank,
					lines: file_lines.lines,
					files: 1,
				})
			} else {
				stats[does_contain].blank += file_lines.blank
				stats[does_contain].lines += file_lines.lines
				stats[does_contain].files += 1
			}
		}
	}

	return stats
}

func main() {
	start := time.Now()

	items := get_files(".")
	t_print(items, stats)

   elapsed := time.Since(start)
	fmt.Printf("Time: %s", elapsed)
}
