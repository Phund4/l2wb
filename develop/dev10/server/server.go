package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Launching server...")

	// Запуск сервера
	serv, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error listening: %v\n", err)
		return
	}
	defer serv.Close()

	for {
		fmt.Println("Waiting for connection...")

		// Подключение к каналу
		conn, err := serv.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error accepting: %v\n", err)
			return
		}
		fmt.Println("New connection. Waiting for messages...")

		go func(conn net.Conn) { // Обрабатываем в горутине для возможности общаться с большим кол-вом клиентов
			defer conn.Close()

			fmt.Printf("Serving new conn %v\n", conn)

			connReader := bufio.NewReader(conn) // Ридер создается один раз

			for {
				// Чтение сообщения
				message, err := connReader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						fmt.Printf("Connection %v closed.\n", conn)
						break
					}
					fmt.Fprintf(os.Stderr, "error reading from conn: %v\n", err)
					break
				}
				message = strings.TrimSpace(message) // Удаляем лишние символы

				fmt.Printf("From: %v Received: %s\n", conn, string(message))

				// Отправка нового сообщения обратно клиенту
				newmessage := strings.ToUpper(message)
				_, err = conn.Write([]byte(newmessage + "\n"))
				if err != nil {
					fmt.Fprintf(os.Stderr, "error writing to conn: %v\n", err)
					break
				}
			}

			fmt.Printf("Done serving client %v\n", conn)
		}(conn)
	}
}