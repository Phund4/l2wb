package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: ", err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, "Error: ", err)
		}

	}
}

func execInput(input string) error {
	//Убираем пробелы и инпуты во входной строке
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)
	//Сплитим на аргументы
	args := strings.Split(input, " ")
	//Соответствие команд задания и виндовс
	switch args[0] {
	case "cd":
		out, err := exec.Command("sh", "-c", fmt.Sprintf("cd %s", args[1])).Output();

		if err != nil {
			return err
		} else {
			fmt.Printf("%s\n", out)
		}
	case "pwd":
		out, err := exec.Command("pwd").Output()

		if err != nil {
			return err
		} else {
			fmt.Printf("%s\n", out)
		}
	case "echo":
		var echoArgs string
		for i := 1; i < len(args); i++ {
			echoArgs = echoArgs + " " + args[i]
		}

		out, err := exec.Command("cmd", "/C", "echo", echoArgs).Output()

		if err != nil {
			return err
		} else {
			fmt.Printf("%s\n", out)
		}
	case "kill":
		out, err := exec.Command("kill", args[1]).Output()

		if err != nil {
			return err
		} else {
			fmt.Printf("%s\n", out)
		}
	case "ps":
		out, err := exec.Command("ps").Output()

		if err != nil {
			return err
		} else {
			fmt.Printf("%s\n", out)
		}
	case "quit":
		os.Exit(0)
	}

	return nil
}
