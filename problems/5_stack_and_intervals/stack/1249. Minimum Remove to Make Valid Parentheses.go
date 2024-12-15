package stack

import "strings"

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

// time: O(n) - проходим всю строку
// mem: O(n) - длина стека может быть равна длине строки
func minRemoveToMakeValid(s string) string {
	// в go строки не изменяемы поэтому нужно сделать список из символов
	// который уже можно менять
	result := []rune(s)
	var stack []int // храним индексы для символа "("

	for i := range result {
		char := result[i]
		if char == '(' {
			stack = append(stack, i)
		} else if char == ')' && len(stack) == 0 {
			// скобка ")" лишняя и должна быть удалена
			result[i] = 0
		} else if char == ')' && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}

	// проходимся по всем лишним скобкам "(" и удаляем их
	for _, i := range stack {
		result[i] = 0
	}

	// делаем строку из элементов списка
	var res strings.Builder
	for _, r := range result {
		if r != 0 {
			res.WriteRune(r)
		}
	}

	return res.String()
}
