package main

import (
	"fmt"
)

func solFunc(s string) string {
	if len(s) == 0 || (s[0] >= '0' && s[0] <= '9') {
		return ""
	}

	var res string

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if i > 0 && s[i-1] >= '0' && s[i-1] <= '9' {
				return ""
			}

			digit := int(s[i] - '0')

			if len(res) == 0 {
				return ""
			}

			lastChar := res[len(res)-1]
			res = res[:len(res)-1]

			if digit > 0 {
				for j := 0; j < digit; j++ {
					res += string(lastChar)
				}
			}

		} else {
			res += string(s[i])
		}
	}

	return res
}

func main() {
	tests := []string{
		"a4bc2d5e",
		"abcd",
		"3abc",
		"45",
		"aaa10b",
		"aaa0b",
		"",
		"d\n5abc",
	}

	for _, t := range tests {
		res := solFunc(t)
		fmt.Printf("Input: %q -> %q\n", t, res)
	}
}
