package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Инициализация флагов
	core := NewCore()
	flag.IntVar(&core.After, "A", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&core.Before, "B", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&core.Context, "C", 0, "'context' (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&core.Count, "c", false, "'count' (количество строк)")
	flag.BoolVar(&core.IgnoreCase, "i", false, "'ignore-case' (игнорировать регистр)")
	flag.BoolVar(&core.Invert, "v", false, "'invert' (вместо совпадения, исключать)")
	flag.BoolVar(&core.Fixed, "F", false, "'fixed', точное совпадение со строкой")
	flag.BoolVar(&core.LineNum, "n", false, "'line num', печатать номер строки")
	flag.Parse()

	// Обновляем параметры after и before в зависимости от параметра context
	core.SyncOutLength()

	// Собираем аргументы
	args := flag.Args()

	// Если аргументов не хватает, то пишем help
	if len(args) < 2 {
		log.Fatalln("Чтобы начать поиск: [флаги] [искомая строка] [название файла]")
	}

	// Получаем искомую фразу
	slicePhrase := args[:len(args)-1]
	// Объединяем в одну строку
	core.Phrase = strings.Join(slicePhrase, " ")

	fileName := args[len(args)-1]

	// Открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Считываем файл
	fileContent, err := io.ReadAll(file);
	if err != nil {
		log.Fatalln(err)
	}

	// Сплит файла
	splitString := strings.Split(string(fileContent), "\n")

	// Используем grep, записываем и печатаем результат
	result := Grep(splitString, core)
	printRes(core, result)
}

// Grep функция поиска фразы или строки в файле с применением доп.условий
func Grep(text []string, c *Core) []*GrepStruct {
	var result []*GrepStruct // Слайс с результатами поиска
	var condition bool // Условие сравнения

	//Проходим построчно по файлу
	for index, str := range text {
		// Если -i, убираем регистр
		if c.IgnoreCase {
			str = strings.ToLower(str)
			c.PhraseToLower()
		}

		// Проверяем условия
		if c.Fixed {
			condition = (strings.Compare(c.Phrase, str) == 0) // полное совпадение строки
		} else {
			condition = strings.Contains(str, c.Phrase) // совпадение подстроки
		}

		// Флаг исключения
		if c.Invert {
			condition = !condition
		}

		// Создаем объект grep
		match := NewGrep()
		// Если выполняется условие, то записываем в эту строку
		if condition {
			c.AddMatch()

			// Определяем количество строк для печати в зависимости от флагов
			var upRange, downRange = 0, len(text) - 1
			if d := index - c.Before; d > upRange {
				upRange = d
			}
			if d := index + c.After; d < downRange {
				downRange = d
			}
			for i := upRange; i <= downRange; i++ {
				match.Result = append(match.Result, Node{
					Key:   i + 1,
					Value: text[i],
				})
			}
			result = append(result, match)
		}

	}
	return result
}

//вывод результата
func printRes(c *Core, res []*GrepStruct) {
	//Если установлен флаг на количество вхождений
	if c.Count {
		fmt.Printf("Совпадений: %d\n", c.CountMatch)
	}

	//Проходим по результату
	for _, match := range res {
		match.Print(c.LineNum)
	}
}