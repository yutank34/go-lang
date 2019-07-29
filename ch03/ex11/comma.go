package ex11

import "strings"

// commaは負でない10進数表記整数文字列にカンマを挿入する
func commaHandler(s string) string {
	if strings.Contains(s, ".") {
		s = floatComma(s)
	} else {
		s = comma(s)
	}
	return s
}

func floatComma(s string) string {
	slice := strings.Split(s, ".")
	return comma(slice[0]) + "." + slice[1]
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + comma(s[n-3:])
}
