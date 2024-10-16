package hashMap

// https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	st := make(map[byte]byte, len(s)) // ключ символ из s, val = t[i]
	ts := make(map[byte]byte, len(t)) // ключ символ из t, val = s[i]

	for i, r := range s {
		if tr, ok := st[byte(r)]; ok && t[i] != tr { // если существует, то проверяем что меняем на тот же символ в t
			return false
		}
		st[byte(r)] = t[i]

		if sr, ok := ts[t[i]]; ok && byte(r) != sr { // если существует, то проверяем что соответствует тому же символу в t
			return false
		}
		ts[t[i]] = byte(r)
	}

	return true
}

/*
Оценка по времени: О(n)
где n - длина строк потому что проходим по строке 1 раз

Оценка по памяти: О(n)
где n длина строк (2n, но 2 в уме)

Описание решения.
Так как длина строк одинакова, то проходимся по 1 строке через for.
Будем добавлять в мапу соответствие буквы из S буквой в T и так же для T к S.
Одновременно проверяем что мапинг из первой соответствует маппингу из второй и наоборот.
Добавляем в мапу после прохождения проверки (та же буква (перезаписываем) или нет в мапе).
Если что-то из этого не верно то возвращаем false.
*/
