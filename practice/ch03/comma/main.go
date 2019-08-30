package comma

// commaは負でない10進数表記整数文字列にカンマを挿入する
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + comma(s[n-3:])
}
