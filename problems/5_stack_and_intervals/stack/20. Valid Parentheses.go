package stack

/*
Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false
*/
// time O(n) - полностью проходим строку
// mem O(n) - длина стека зависит от длины строки
func isValid(s string) bool {
	var stack []byte
	pairs := map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	for _, char := range s {
		if _, ok := pairs[byte(char)]; ok { // добавляем в стек отрытую скобку
			stack = append(stack, byte(char))
			continue
		}

		if len(stack) == 0 { // если скобка сразу же закрыта
			return false
		}

		lastChar := stack[len(stack)-1] // удаляем последний элемент из стека
		stack = stack[:len(stack)-1]
		// проверяем что последний элемент в стеке и текущий образуют пару
		// если пару не образуют, то вернем False
		if pairs[lastChar] != byte(char) {
			return false
		}
	}

	return len(stack) == 0
}
