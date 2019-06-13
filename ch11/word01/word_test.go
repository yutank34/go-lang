package word

import "testing"

func TestPalindrome(t *testing.T) {
	// if !IsPalindrome("detartrated") {
	// 	t.Error(`IsPalindrome("detartrated") = false`)
	// }
	// if !IsPalindrome("kayak") {
	// 	t.Error(`IsPalindrome("kayak") = false`)
	// }
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"ätä", true},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("ätä") {
		t.Error(`IsPalindrome("ätä") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome("%q") = false`, input)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("kayak") = true`)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}
