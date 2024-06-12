package main

import (
	"fmt"
	"sync"
	"time"
)

// Store of events
type Store struct {
	mu     *sync.Mutex
	events map[int][]Event
}

// Create event
func (s *Store) Create(e *Event) error {

	// Блокируем структуру мьютексом
	s.mu.Lock()
	defer s.mu.Unlock()

	if events, ok := s.events[e.UserID]; ok {
		for _, event := range events {
			if event.EventID == e.EventID {
				return fmt.Errorf("event with such id (%v) already present for this user (%v);", e.EventID, e.UserID)
			}
		}
	}

	s.events[e.UserID] = append(s.events[e.UserID], *e)

	return nil
}

// Update event
func (s *Store) Update(e *Event) error {

	// Блокируем структуру мьютексом
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1

	// Проверка наличия события с нужным userID
	if _, ok := s.events[e.UserID]; !ok {
		return fmt.Errorf("user with this id (%v) already exists", e.UserID)
	}

	// Если пользователь есть, то ищем именно его
	for idx, event := range s.events[e.UserID] {
		if event.EventID == e.EventID {
			index = idx
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("user with id = (%v) doesn't have event with id = (%v)", e.UserID, e.EventID)
	}

	s.events[e.UserID][index] = *e

	return nil
}

// Delete event
func (s *Store) Delete(e *Event) (*Event, error) {

	// Блокируем структуру мьютексом
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1

	// Проверяем наличие пользователя с подобным userID
	if _, ok := s.events[e.UserID]; !ok {
		return nil, fmt.Errorf("user with this id (%v) doesn't exist", e.UserID)
	}

	// Проверяем наличие события с подобным id
	for idx, event := range s.events[e.UserID] {
		if event.EventID == e.EventID {
			index = idx
			break
		}
	}

	if index == -1 {
		return nil, fmt.Errorf("user with id = (%v) doesn't have event with id = (%v)", e.UserID, e.EventID)
	}

	eventsLength := len(s.events[e.UserID])
	deletedEvent := s.events[e.UserID][index]
	s.events[e.UserID][index] = s.events[e.UserID][eventsLength-1]
	s.events[e.UserID] = s.events[e.UserID][:eventsLength-1]

	return &deletedEvent, nil
}

// GetEventsForDay - Получить события за конкретный день
func (s *Store) GetEventsForDay(userID int, date time.Time) ([]Event, error) {

	// Блокируем структуру мьютексом
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	// Проверяем наличие пользователя с подобным userID
	if _, ok := s.events[userID]; !ok {
		return nil, fmt.Errorf("user with this id (%v) doesn't exist", userID)
	}

	// Поиск событий
	for _, event := range s.events[userID] {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}

	return result, nil
}

// GetEventsForWeek - Получить события за неделю
func (s *Store) GetEventsForWeek(userID int, date time.Time) ([]Event, error) {

	// Блокируем структуру мьютексом
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	// Проверяем наличие пользователя с подобным userID
	if _, ok := s.events[userID]; !ok {
		return nil, fmt.Errorf("user with this id (%v) doesn't exist", userID)
	}

	// Поиск событий
	for _, event := range s.events[userID] {
		y1, w1 := event.Date.ISOWeek()
		y2, w2 := date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			result = append(result, event)
		}
	}

	return result, nil
}

// GetEventsForMonth - Получить события за месяц
func (s *Store) GetEventsForMonth(userID int, date time.Time) ([]Event, error) {

	// Блокируем структуру мьютексом
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	// Проверяем наличие пользователя с подобным userID
	if _, ok := s.events[userID]; !ok {
		return nil, fmt.Errorf("user with this id (%v) doesn't exist", userID)
	}

	// Поиск событий
	for _, event := range s.events[userID] {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}

	return result, nil
}