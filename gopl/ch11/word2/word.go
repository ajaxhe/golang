package word

import "unicode"

/*
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
*/

func IsPalindrome(s string) bool {
	n := len(s)
	letters := make([]rune, 0, n)
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	/*
		c := utf8.RuneCountInString(s)
		// utf8编码的字符串可能会有问题
		for i := 0; i < c/2; i++ {
			if letters[i] != letters[c-1-i] {
				return false
			}
		}
	*/
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
