package main

import (
	"fmt"
	"regexp"
)

var text = "my email is 393673587hh@qq.com" +
	" abc@163.org. Email2 is dfsff@hh.cn!"

func main() {
	re  := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)

	for _, val := range match {
		fmt.Println(val)
	}
}