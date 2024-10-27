package hashMap

// https://leetcode.com/problems/palindrome-permutation/description/
// Given a string 's' return true if a permutation of the string could form a palindrome and false otherwise.

func canPermutePalindrome(s string) bool {
	counter := [26]rune{}

	for _, symbol := range s {
		counter[symbol-'a'] = (counter[symbol-'a'] + 1) % 2
	}

	isOdd := len(s)%2 != 0
	oddCounter := 0

	for i := range counter {
		if counter[i]%2 != 0 {
			oddCounter++
		}
	}

	return (isOdd && oddCounter == 1) || oddCounter == 0
}

/*
Оценка по времени: О(n)
Проходим строку и заполняем массив и потом еще раз проходим по массиву - 2 * n

Оценка по памяти: О(n)
Выделяем памяти для массива как n (длина строки)

Описание решения
Проходим по строке, заполняем массив значениями 1 или 0
Далее вводими счетчик количества букв с нечетными значениями.
Если длина строки тоже нечетная, то возможна 1 такая буква, если четная, то 0
*/
