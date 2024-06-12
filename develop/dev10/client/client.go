package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	exitChan := make(chan os.Signal, 1)
	// При нажатии ctrl + D в канал sigCh будет отправлено сообщение
	signal.Notify(exitChan, syscall.SIGQUIT)
	go Exit(exitChan)

	// Устанавливаем флаги на задержку и аргументы на хост и порт
	timeout := flag.String("timeout", "10s", "timeout for a connection")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Please enter timeout")
		return
	}

	// Переводим задержку в нужный формат
	timeoutDuration, err := time.ParseDuration(*timeout)
	if err != nil {
		fmt.Println("Error parsing timeout")
		return
	}

	// Формируем строку подключения
	host := flag.Arg(0)
	port := flag.Arg(1)
	hostPort := host + ":" + port

	// Подключаемся к сокету. Задаем таймаут подключения
	conn, err := net.DialTimeout("tcp", hostPort, timeoutDuration)
	if err != nil {
		fmt.Println("Error connection to", hostPort)
		return
	}
	defer conn.Close()

	// Создаем ридеры
	console := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	for {
		fmt.Print("Your message: ")

		// Чтение сообщения с консоли
		text, err := console.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading message: %v", err)
			return
		}
		text = strings.TrimSpace(text) // Удаляем лишние символы

		// Отправляем сообщение
		fmt.Fprintf(conn, text+"\n")
		if text == "exit" {
			fmt.Println("Closing connection...")
			return
		}

		// Получаем ответ
		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response from server: %v", err)
			return
		}
	
		// Удаляем лишние символы
		message = strings.TrimSpace(message)
		fmt.Printf("From server: %s\n", message)
	}
}

// Ждем нужного сигнала для выхода из программы
func Exit(exitChan chan os.Signal) {
	for {
		switch <-exitChan {
		case syscall.SIGQUIT:
			fmt.Println("Press ctrl + D to exit")
			os.Exit(0)
		default:
		}
	}
}