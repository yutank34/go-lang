package basename1

// basenameはディレクトリ要素と接尾辞を取り除く
// e.g., a=> a, a.go => a, a/b/c.go => x, a/b.c.go => b.c
func basename(s string) string {
	// 最後の'/'とその前の全てを破棄する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 最後の'.'より前の情報を全て保持する
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
