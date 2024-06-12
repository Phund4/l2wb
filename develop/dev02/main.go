package main;

import (
	"fmt"
)

func main() {
	fmt.Println(extract("45"))
}

func extract(str string) string {

	// Необходимые предусловия
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return str
	}

	var res []rune

	letters := []rune(str)

	if isNumber(letters[0]) {
		return "Incorrect string"
	}

	// Перебор строки
	for len(letters) > 0 {
		// Если последний оставшийся символ число, то возвращаем ошибку
		if len(letters) == 1 {
			if isNumber(letters[0]) {
				return "Incorrect string"
			}
			res = append(res, letters[0])
			break
		}

		// Смотрим, чем являются два первых символа
		a, b := letters[0], letters[1]

		// Если оба число, то возвращаем ошибку
		if isNumber(a) && isNumber(b) {
			return "Incorrect string"
		}

		// Если второй символ число, то добавляем такое количество букв
		if isNumber(b) {
			res = append(res, repeatedLetter(int(b-'0'), a)...) // int(b-'0') преобразование rune в int
			letters = letters[2:]
		} else {
			// Если второй символ не число, то добавляем новую букву в массив
			res = append(res, a)
			letters = letters[1:]
		}
	}
	
	return string(res)
}

// Проверка символа на число
func isNumber(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

// Функция создания массива заданной длины из заданного символа
func repeatedLetter(n int, symb rune) []rune {
	res := make([]rune, n);

	for i := 0; i < n; i++ {
		res = append(res, symb)
	}

	return res
}