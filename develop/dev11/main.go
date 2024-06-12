package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

// Используемая дата
const dateLayout = "2006-01-02"

// Глобальная переменная для хранения событий
var storage Store = Store{events: make(map[int][]Event), mu: &sync.Mutex{}}

func main() {
	mux := http.NewServeMux()

	// POST
	mux.HandleFunc("/create_event", createEventHandler)
	mux.HandleFunc("/update_event", updateEventHandler)
	mux.HandleFunc("/delete_event", deleteEventHandler)

	// GET
	mux.HandleFunc("/events_for_day", eventsForDayHandler)
	mux.HandleFunc("/events_for_week", eventsForWeekHandler)
	mux.HandleFunc("/events_for_month", eventsForMonthHandler)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	// Чтение порта из .env
	port := ":8080"
	func() {
		temp := os.Getenv("PORT")
		if temp != "" {
			port = temp
		}
	}()

	// Реализация middleware
	wrappedMux := NewLogger(mux)

	log.Fatalln(http.ListenAndServe(":"+port, wrappedMux))
}

// Обработчик создания события
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	// Получаем данные запроса и вносим в структуру события
	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидируем поля
	if err := e.Validate(); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Записываем событие в глобальную переменную
	if err := storage.Create(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Event was create successful", []Event{e}, http.StatusCreated)

	fmt.Println(storage.events)
}

// Обработчик обновления события
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	// Получаем данные
	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидируем поля
	if err := e.Validate(); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Обновляем событие
	if err := storage.Update(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Event was update successful", []Event{e}, http.StatusOK)

	fmt.Println(storage.events)
}

// Обработчик удаления события
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	// Получаем данные
	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var deletedEvent *Event
	var err error
	// Удаляем событие
	if deletedEvent, err = storage.Delete(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Event was delete successful", []Event{*deletedEvent}, http.StatusOK)

	fmt.Println(storage.events)
}

// Получить события за конкретный день
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем userID
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем date
	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем среди событий необходимые
	var events []Event
	if events, err = storage.GetEventsForDay(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request was complete successful", events, http.StatusOK)
}

// Получить события за конкретную неделю
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем userID
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем date
	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем среди событий необходимые
	var events []Event
	if events, err = storage.GetEventsForWeek(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request was executed successful", events, http.StatusOK)
}

// Получить события за конкретный год
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем userID
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получаем date
	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем среди событий необходимые
	var events []Event
	if events, err = storage.GetEventsForMonth(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request was executed successful", events, http.StatusOK)
}

// Ответ с ошибкой
func errorResponse(w http.ResponseWriter, e string, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{Error: e}

	js, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Ответ без ошибки
func resultResponse(w http.ResponseWriter, r string, e []Event, status int) {
	resultResponse := struct {
		Result string  `json:"result"`
		Events []Event `json:"events"`
	}{Result: r, Events: e}

	js, err := json.Marshal(resultResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}