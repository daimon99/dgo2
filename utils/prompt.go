package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DATE_FORMAT = "2006-01-02"

func Test() {
	print("just a test")
}

func defaultInput(prompt_text string, default_input string) {
	if len(default_input) != 0 {
		prompt_text = fmt.Sprintf("%s [%s]: ", prompt_text, default_input)
	}
	fmt.Printf("%s: ", prompt_text)
}

func Prompt(prompt_text string, default_input string) (ret string) {
	defaultInput(prompt_text, default_input)
	var user_input string
	fmt.Scanln(&user_input)
	if len(user_input) == 0 {
		return default_input
	}
	return user_input
}

func PromptInt(prompt_text string, default_input int) (ret int) {
RE:
	user_input := Prompt(prompt_text, strconv.Itoa(default_input))
	ret, err := strconv.Atoi(user_input)
	if err != nil {
		fmt.Println("请输入一个整数")
		goto RE
	}
	return ret
}

func PromptDate(prompt_text string, default_input string) (ret time.Time) {
RE:
	user_input := Prompt(prompt_text, default_input)
	on_date, err := time.Parse(DATE_FORMAT, user_input)
	if err != nil {
		fmt.Println("日期格式不对，必须是 2006-01-02 格式")
		goto RE
	}
	return on_date
}

func PromptBool(prompt_text string, default_input bool) (ret bool) {
RE:
	default_input_str := "n"
	if default_input {
		default_input_str = "y"
	}
	user_input := strings.ToLower(Prompt(prompt_text, default_input_str))
	switch user_input {
	case "y", "true", "yes":
		return true
	case "n", "false", "no":
		return false
	default:
		fmt.Println("输入非法，请重新输入")
		goto RE
	}

}
