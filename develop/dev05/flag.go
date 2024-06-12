package main

import "strings"

// Core - Структура флагов
type Core struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Phrase     string
	CountMatch int
}

// NewCore - Инициализация Core'a
func NewCore() *Core {
	return &Core{
		After:      0,
		Before:     0,
		Context:    0,
		Count:      false,
		IgnoreCase: false,
		Invert:     false,
		Fixed:      false,
		LineNum:    false,
		Phrase:     "",
		CountMatch: 0,
	}
}

// PhraseToLower - Перевод искомой строки в нижний регистр
func (c *Core) PhraseToLower() {
	c.Phrase = strings.ToLower(c.Phrase)
}

// SyncOutLength - Обновляем параметры after и before в зависимости от параметра context
func (c *Core) SyncOutLength() {
	if c.Context > c.After {
		c.After = c.Context
	}
	if c.Context > c.Before {
		c.Before = c.Context
	}
}

// AddMatch - Инкрементируем количество совпадений
func (c *Core) AddMatch() {
	c.CountMatch++
}