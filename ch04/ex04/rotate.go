package ex04

func rotate(s []int, step int) {
	if step == 0 {
		return
	}
	if len(s) < step {
		step = step - len(s)
	}
	tmp := append(s, s[:step]...)
	tmp = append(tmp[step:])
	for i := 0; i < len(s); i++ {
		s[i] = tmp[i]
	}
}

// func rotate(s []int) {
// 	reverse(s[:2])
// 	reverse(s[2:])
// 	reverse(s)
// }

// func reverse(s []int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }
