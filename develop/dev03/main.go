package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type indexValue struct {
	index int
	value string
}

func main() {
	var column int
	var n, r, u bool
	flag.IntVar(&column, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	var input io.Reader

	// Проверяем указано ли имя файла
	// Если все ок то считываем файл и обрабатываем возможные ошибки
	if filename := flag.Arg(0); filename == "" {
		fmt.Printf("Укажите имя файла!\n")
		os.Exit(1)
	} else {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %s", err)
			os.Exit(1)
		}
		defer f.Close()

		input = f
	}

	var lines [][]string

	//Создаем матрицу из строк, где каждая строка это отдельный элемент
	buf := bufio.NewScanner(input)
	for buf.Scan() {
		line := buf.Text()
		lines = append(lines, strings.Split(line, " "))
	}

	//Создаем массив из indexValue и кладем туда номер строки и саму строку
	var stringsForSort []indexValue
	for index, lineArr := range lines {
		stringsForSort = append(stringsForSort, indexValue{
			index: index,
			value: lineArr[column],
		})
	}

	// Создаем файл с результатом работы
	file, err := os.Create("output.txt");
	if err != nil {
		fmt.Printf("Ошибка создания файла: %s", err)
		os.Exit(1)
	}
	defer file.Close();

	//Теперь, имея массив из элементов для сортировки, применим функцию сортировки с нужными флагами
	//Получаем индексы строк и выводим их
	sortedLinesOrder := sortIndexValueArray(stringsForSort, r)
	for i, val := range sortedLinesOrder {
		var dismiss bool

		//проверяем флаг на удаление дубликатов
		if u {
			//не выходиим за диапазон
			if len(lines) > i+1 {
				//сравниваем элементы этой и следующей строки, если совпадают то пропускаем одну, не печатаем
				if len(lines[i]) == len(lines[i+1]) {
					for id := range lines[i] {
						if lines[i][id] != lines[i+1][id] {
							dismiss = true
						}
					}
				}
			} else {
				fmt.Println(lines[val])
				continue
			}

			if !dismiss {
				continue
			}
		}

		resString := strings.Join(lines[val], " ");
		file.WriteString(resString + "\n")
	}
}

//Принимаем на вход массив с данными которые нужно сортировать и флаги
func sortIndexValueArray(stringsForSort []indexValue, r bool) []int {
	index := make(map[string]int)
	var sorting []string

	//Создаем мапу в которой сортруемое слово это ключ, а номер строки это значение, а также просто слайс из значений
	for _, val := range stringsForSort {
		index[val.value] = val.index
		sorting = append(sorting, val.value)
	}

	//сортируем слайс
	sort.Strings(sorting)

	//Для итогового результата
	var res []int

	//Если флаг false, то выводим в порядке возрастания
	//У нас есть отсортированный слайс, поэтому в результате записываем соответствие номера строки и ключа в новой сортировке
	switch r {
	case false:
		for i := 0; i < len(sorting); i++ {
			res = append(res, index[sorting[i]])
		}
	case true:
		for i := len(sorting) - 1; i >= 0; i-- {
			res = append(res, index[sorting[i]])
		}
	}

	return res
}