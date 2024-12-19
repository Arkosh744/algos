package stack

import "strconv"

/*
Вход: tokens = ["2", "1", "+", "3", "*"]
Пошаговая обработка:
"2" -> добавляем в стек -> stack = [2]
"1" -> добавляем в стек -> stack = [2, 1]
"+" -> выполняем 2 + 1 -> stack = [3]
"3" -> добавляем в стек -> stack = [3, 3]
"*" -> выполняем 3 * 3 -> stack = [9]
Результат: 9
*/

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		res := 0
		switch tokens[i] {
		case "+":
			res = stack[len(stack)-2] + stack[len(stack)-1]
		case "-":
			res = stack[len(stack)-2] - stack[len(stack)-1]
		case "*":
			res = stack[len(stack)-2] * stack[len(stack)-1]
		case "/":
			res = stack[len(stack)-2] / stack[len(stack)-1]
		default:
			v, _ := strconv.Atoi(tokens[i])
			stack = append(stack, v)
			continue
		}

		stack = stack[:len(stack)-2]
		stack = append(stack, res)
	}

	return stack[len(stack)-1]
}
