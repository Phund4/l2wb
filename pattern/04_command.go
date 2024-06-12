package pattern;

import (
	"fmt"
)

// Button struct
type Button struct {
	command Command
}

// Метод выполнения процесса при нажатии на кнопку
func (b *Button) press() {
	b.command.execute();
}

// Command interface
type Command interface {
	execute()
}

// OnCommand struct
type OnCommand struct {
	device Device
}

// Основной метод для конкретной команды
func (on *OnCommand) execute() {
	on.device.on()
}

// OffCommand struct
type OffCommand struct {
	device Device
}

// Основной метод для конкретной команды
func (off *OffCommand) execute() {
	off.device.off()
}

// Device interface
type Device interface {
	on()
	off()
}

// TV struct
type TV struct {
	isRunning bool
}

// Метод включения TV
func (tv *TV) on() {
	tv.isRunning = true;
	fmt.Println("Turning TV on")
}

// Метод выключения TV
func (tv *TV) off() {
	tv.isRunning = false;
	fmt.Println("Turning TV off")
}

// CommandUserCode Пользовательский код использования паттерна команда
func CommandUserCode() {
	tv := &TV{};

	onCommand := OnCommand{
		device: tv,
	}

	offCommand := OffCommand{
		device: tv,
	}

	onButton := Button{
		command: &onCommand,
	}
	onButton.press();

	offButton := Button{
		command: &offCommand,
	}
	offButton.press();
}

/*
Паттерн команда
Применимость: 1) Когда нам нужно использовать операцию как объект. А объект можно хранить и передавать.
2) Когда мы хотим ставить операции в очередь, выполнять их по расписанию или передавать по сети.
Плюсы: 1) Убирает прямую зависимость от получателей и клиентов
2) Позволяет реализовать отмену и повторной выполнение операций
3) Позволяет собирать сложные команды из простых
4) Реализует принцип открытости/закрытости
Минусы: 1) Усложняет код программы из-за большого кол-ва дополнительных классов
Пример использования: В каком нибудь сервисе есть несколько кнопок, выполняющие 
одно и тоже действие. Это взаимодействие с бизнес-логикой, которое необходимо
отделить и не засорять код одними и теми же реализациями действий.
*/