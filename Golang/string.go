package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	str = "Yes,我爱我的国!"

	fmt.Println(str)
	for _, ch := range []byte(str) {
		fmt.Printf("%x ", ch) //59 65 73 2c e6 88 91 e7 88 b1 e6 88 91 e7 9a 84 e5 9b bd 21
	}
	fmt.Println()

	for index, ch := range str { //ch is a rune （int32）
		fmt.Printf("(%d, %x)", index, ch) //(0, 59)(1, 65)(2, 73)(3, 2c)(4, 6211)(7, 7231)(10, 6211)(13, 7684)(16, 56fd)(19, 21)
	}
	fmt.Println()
	fmt.Println(len(str)) //20

	fmt.Println("Rune count：", utf8.RuneCountInString(str)) //10

	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	} //Yes,我爱我的国!
	fmt.Println()

	for index, ch := range []rune(str) {
		fmt.Printf("(%d, %c)", index, ch) //(0, Y)(1, e)(2, s)(3, ,)(4, 我)(5, 爱)(6, 我)(7, 的)(8, 国)(9, !)
	}
}
