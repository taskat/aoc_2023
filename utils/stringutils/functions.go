package stringutils

import "strings"

func Replace(s string, old string, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func ReplaceFromBack(s string, old string, new string, n int) string {
	reversedS := Reverse(s)
	reversedOld := Reverse(old)
	reversedNew := Reverse(new)
	result := strings.Replace(reversedS, reversedOld, reversedNew, n)
	return Reverse(result)
}

func Reverse(s string) string {
	newString := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		newString[i] = s[len(s)-i-1]
	}
	return string(newString)
}
