package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Функция создает канал, который закрывается спустя заданное время
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()

	// or получает список каналов и вернет первый завершившийся
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Done after %v\n", time.Since(start))
}

func or(ch ...<-chan interface{}) <-chan interface{} {
	// Создаем single канал
	single := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем горутины, прослушивающие каналы из списка
	for _, channel := range ch {
		// Анонимная функция, смотрит за каналом, ожидая его завершения
		go func(channel <-chan interface{}) {
			// В каждом запускаем цикл, который завершится по закрытии канала
			for range channel {}
			wg.Done()
		}(channel)
	}

	wg.Wait()

	// Закрываем и возвращаем single канал
	close(single)
	return single
}
