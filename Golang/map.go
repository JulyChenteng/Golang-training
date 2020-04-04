package main

import "fmt"

/*寻找最长不含有重复字符的子串*/
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start, maxLength := 0, 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch]
		}

		if i-start > maxLength {
			maxLength = i - start
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
}
