package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Event struct
type Event struct {
	UserID      int       `json:"user_id"`
	EventID     int       `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Decode - Десереализация данных
func (e *Event) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Validate - Проверка на правильное заполнение полей
func (e *Event) Validate() error {
	if e.UserID <= 0 {
		return fmt.Errorf("incorrect user_id: %v;", e.UserID)
	}

	if e.EventID <= 0 {
		return fmt.Errorf("incorrect event_id: %v;", e.EventID)
	}

	if e.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	return nil
}