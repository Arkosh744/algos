package window

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// time: O(n)
// mem:  O(k), где k - число уникальных символов в строке
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		result = 0
		l, r   = 0, 0
	)

	windowChars := map[rune]bool{rune(s[0]): true}
	for l < len(s) {
		for r+1 < len(s) && !windowChars[rune(s[r+1])] {
			windowChars[rune(s[r+1])] = true
			r++
		}

		result = max(result, r-l+1)
		delete(windowChars, rune(s[l])) // когда двигаем левую границу убираем буквы
		l++

		// если r < l то обновляем r и windowChars
		if r < l && l < len(s) {
			r = l
			windowChars = map[rune]bool{rune(s[r]): true}
		}
	}

	return result
}

/*

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
