package ex10

import "bytes"

// commaは負でない10進数表記整数文字列にカンマを挿入する
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	// 100000 => 100,
	pref := n % 3
	if pref == 0 {
		pref = 3
	}
	buf.WriteString(s[:pref])
	for i := 1; i <= n/3; i++ {
		if pref+1*3 > n {
			break
		}
		buf.WriteString(",")
		buf.WriteString(s[pref+(i-1)*3 : pref+i*3])
	}
	return buf.String()
}
