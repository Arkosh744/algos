package pointers

func backspaceCompare(s string, t string) bool {
	p1, p2 := len(s), len(t)
	for p1 > 0 && p2 > 0 { // двигаемся справа налево <---
		p1 = findNextSymbol(s, p1-1) // если не удаляем, то смещение на 1, либо считаем удаления
		p2 = findNextSymbol(t, p2-1)
		if p1 >= 0 && p2 >= 0 && s[p1] != t[p2] { // сравниваем буквы на текущем месте
			return false
		}
	}

	return findNextSymbol(s, p1-1) == findNextSymbol(t, p2-1) // какой-то из массивов может закончится быстрее, не нужно полное сравнение

}

func findNextSymbol(s string, i int) int {
	toRemove := 0
	for i >= 0 && (toRemove > 0 || s[i] == '#') {
		if s[i] == '#' { // скипаем удаление
			toRemove += 1
			i -= 1
			continue
		}

		toRemove -= 1 // "удаляем" символ и смещаемся
		i -= 1
	}

	return i
}

/*
https://leetcode.com/problems/backspace-string-compare/description/
Example 1:

Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Example 2:
Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

Example 3:
Input: s = "a#c", t = "b"
Output: false
Explanation: s becomes "c" while t becomes "b".
*/
