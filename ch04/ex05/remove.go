package ex05

func removeDuplicate(s []string) {
	v := make([]string, len(s), cap(s))
	for i, j := 0, 0; i < len(s); i++ {
		if i == len(s)-1 {
			v[j] = s[i]
			break
		}
		if s[i] == s[i+1] {
			continue
		}
		v[j] = s[i]
		j++
	}

	for i := 0; i < cap(s); i++ {
		s[i] = v[i]
	}
}
