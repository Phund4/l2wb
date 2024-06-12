package pattern

import "fmt"

// IGun interface
type IGun interface {
    setName(name string)
    setPower(power int)
    getName() string
    getPower() int
}

// Gun struct
type Gun struct {
	name string
	power int
}

func (g *Gun) getName() string {
	return g.name;
}

func (g *Gun) getPower() int {
	return g.power;
}

func (g *Gun) setName(name string) {
	g.name = name;
}

func (g *Gun) setPower(power int) {
	g.power = power;
}

// Ak struct
type Ak struct {
	Gun
}

// NewAk Функция создания экземпляра Ak
func NewAk() IGun {
	return &Ak{
		Gun: Gun{
			name: "AK",
			power: 4,
		},
	}
}

// Musket struct
type Musket struct {
	Gun
}

// NewMusket Функция создания экземпляра Musket
func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name: "Musket",
			power: 7,
		},
	}
}

// GetGun Ключевая функция в паттерне, порождающая оружие
func GetGun(gunType string) (IGun, error) {
	if gunType == "ak" {
		return NewAk(), nil;
	} else if gunType == "musket" {
		return NewMusket(), nil;
	}
	return nil, fmt.Errorf("incorrect gun type: %s", gunType);
}

// FactoryMethodUserCode Пользовательский код использования паттерна фабричный метод
func FactoryMethodUserCode() {
	ak47, _ := GetGun("ak47")
    musket, _ := GetGun("musket")

	fmt.Printf("Gun: %s\nPower: %d\n", ak47.getName(), ak47.getPower())
	fmt.Printf("Gun: %s\nPower: %d\n", musket.getName(), musket.getPower())
}

/*
Паттерн фабричный метод
Применимость: 1) Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать наш код
2) Когда мы хотим дать пользователю возможность расширить часть библиотеки/фреймворка
Плюсы: 1) Выделяет код производства продуктов в одно место, упрощая поддержку кода
2) Упрощает добавление новых продуктов в программу
3) Реализует принцип открытости/закрытости
Минусы: 1) Может засориться иерархия классов, т.к. 
для каждого класса продукта нужно создавать подкласс его создателя
Пример использования: допустим пользователь может взаимодействовать с нашим продуктом
через web и desktop версии, тогда необходимо создать две различные кнопки одного функционала
под каждую версию. С этим поможет паттерн фабричный метод.
*/