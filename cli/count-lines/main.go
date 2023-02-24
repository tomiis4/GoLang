package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LinesCount struct {
	language string
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

func t_print(arr []LinesCount) {
	top_str := "Language | Lines | Blank | Comments"

	fmt.Println(repeat(len(top_str), "-"))
	fmt.Println(top_str)
	fmt.Println(repeat(len(top_str), "-"))

	for _, elem := range arr {
		lang_len := len(elem.language)
		lines_len := len(strconv.Itoa(elem.lines))
		blank_len := len(strconv.Itoa(elem.blank))

		lang_spaces := repeat(11-lang_len, " ")
		format_lang_lines := fmt.Sprintf("%s%s%d", elem.language,lang_spaces , elem.lines)

		blank_spaces := repeat(19-(lines_len+lang_len+len(lang_spaces)), " ")
		format_blank := fmt.Sprintf("%s%d", blank_spaces, elem.blank)

		comment_spaces := repeat(27-(blank_len+lines_len+lang_len+len(lang_spaces)+len(blank_spaces)), " ")
		format_command := fmt.Sprintf("%s%d", comment_spaces, elem.comments)

		formated_str := fmt.Sprintf("%s%s%s", format_lang_lines, format_blank, format_command)
		fmt.Println(formated_str)
	}

	fmt.Println(repeat(len(top_str), "-"))
}

func main() {
	var obj LinesCount
	obj.language = "GoLang"
	obj.comments = 51
	obj.blank = 240
	obj.lines = 197

	var obj2 LinesCount
	obj2.language = "CPP"
	obj2.comments = 11
	obj2.blank = 512
	obj2.lines = 19823

	t_print([]LinesCount{obj, obj2})
}
