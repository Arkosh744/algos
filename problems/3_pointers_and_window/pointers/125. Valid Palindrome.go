package pointers

import (
	"unicode"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		// в условии сказано, что у нас ASCII символы, поэтому так делать можно
		leftRune, rightRune := rune(s[l]), rune(s[r])

		if !unicode.IsLetter(leftRune) && !unicode.IsDigit(leftRune) { // скипаем слева не букву
			l++
			continue
		}
		if !unicode.IsLetter(rightRune) && !unicode.IsNumber(rightRune) { // скипаем спарава не букву
			r--
			continue
		}

		if unicode.ToLower(leftRune) != unicode.ToLower(rightRune) { // сравниваем буквы/цифры
			return false
		}
		// сдвигаем поинтеры одновременно
		l++
		r--
	}

	return true
}
