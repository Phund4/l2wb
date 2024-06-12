package main

import (
	"fmt"
	"sort"
)

// Node - Структура узла
type Node struct {
	Key   int
	Value string
}

// GrepStruct - Структура Grep
type GrepStruct struct {
	Result []Node
}

// NewGrep - Инициализация GrepStruct 
func NewGrep() *GrepStruct {
	return &GrepStruct{
		Result: []Node{},
	}
}

// SortResultASC - Сортировка строк результата по возрастанию
func (g *GrepStruct) SortResultASC() {
	sort.Slice(g.Result, func(i, j int) bool {
		return g.Result[i].Key < g.Result[j].Key
	})
}

// Print - Вывод результата
func (g *GrepStruct) Print(withIndex bool) {
	g.SortResultASC()
	switch withIndex {
	// Если нужно печатать номер строки, то этот вариант
	case true:
		for _, v := range g.Result {
			fmt.Printf("%d. %s\n", v.Key, v.Value)
		}
	default:
		for _, v := range g.Result {
			fmt.Printf("%s\n", v.Value)
		}
	}
}
