package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arr := []string{"листок", "пятак", "пятка", "Пятка", "пятка", "слиток", "столик", "тяпка", "кремень", "чямик", "мячик"}
	res := makeAnagram(&arr);
	fmt.Println(res);
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

// Пользовательская сортировка строк
func sortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

// Метод поиска строки в неотсортированном массиве строк
func searchString(str string, arr []string) bool {
	for _, el := range arr {
		if strings.Compare(str, el) == 0 {
			return true;
		}
	}
	return false;
}

// Метод создания словаря анаграмм из массива строк
func makeAnagram(arr *[]string) *map[string][]string {
	res := make(map[string][]string);

	// Словарь для опознания анаграмм
	anagrams := make(map[string]string);

	// Проход по массиву и заполнение словаря
	for _, el := range *arr {
		word := strings.ToLower(el);
		anagram := sortString(word);
		val, ok := anagrams[anagram];
		if !ok {
			anagrams[anagram] = word;
			res[word] = []string{word};
		} else if !searchString(word, res[val]) {
			res[val] = append(res[val], word);
		}
	}

	for key, value := range res {
		// Удаление одиночного элемента
		if len(value) == 1 {
			delete(res, key);
			continue;
		}

		// Сортировка массивов
		if !sort.StringsAreSorted(value) {
			sort.Strings(value)
		}
	}

	return &res;
}
