package pattern

import "fmt"

// IBuilder Интерфейс строителя
type IBuilder interface {
	setWindowType()
	setDoorType()
	getHouse() House
}

// Создание строителя через пользовательский запрос
func getBuilder(builderType string) IBuilder {
	if builderType == "normalBuilder" {
		return NewNormalBuilder()
	} else if builderType == "castleBuilder" {
		return NewCastleBuilder()
	}
	return nil
}

// House struct
type House struct {
	windowType string
	doorType   string
}

// NormalBuilder struct
type NormalBuilder struct {
	windowType string
	doorType   string
}

// NewNormalBuilder Функция создания NormalBuilder
func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

// Метод установки window
func (n *NormalBuilder) setWindowType() {
	n.windowType = "normal window"
}

// Метод установки door
func (n *NormalBuilder) setDoorType() {
	n.doorType = "normal door"
}

// Функция получения объекта дома
func (n *NormalBuilder) getHouse() House {
	return House{
		windowType: n.windowType,
		doorType:   n.doorType,
	}
}

// CastleBuilder struct
type CastleBuilder struct {
	windowType string
	doorType   string
}

// NewCastleBuilder Функция создания CastleBuilder
func NewCastleBuilder() *CastleBuilder {
	return &CastleBuilder{}
}

// Метод установки window
func (c *CastleBuilder) setWindowType() {
	c.windowType = "castle window"
}

// Метод установки door
func (c *CastleBuilder) setDoorType() {
	c.doorType = "castle door"
}

// Функция получения объекта дома
func (c *CastleBuilder) getHouse() House {
	return House{
		windowType: c.windowType,
		doorType:   c.doorType,
	}
}

// Director struct
type Director struct {
	builder IBuilder
}

// NewDirector Функция создания Director
func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// Метод изменения строителя
func (d *Director) setBuilderType(b IBuilder) {
	d.builder = b
}

// Метод постройки дома при помощи строителя
func (d *Director) buildHouse() House {
	d.builder.setWindowType()
	d.builder.setDoorType()
	return d.builder.getHouse()
}

// BuilderUserCode Пользовательский код использования паттерна строитель
func BuilderUserCode() {
	normalBuilder := getBuilder("normal")
    iglooBuilder := getBuilder("igloo")

    director := NewDirector(normalBuilder)
    normalHouse := director.buildHouse()

    fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
    director.setBuilderType(iglooBuilder)
    iglooHouse := director.buildHouse()

    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
}

/*
Паттерн строитель
Применимость: 1) Строитель позволяет создавать объекты опционально, выбирая каждый раз
только необходимые параметры, игнорируя ненужные.
2) Строитель позволяет создавать несколько представлений одного объекта,
отличающиеся в деталях.
Плюсы: 1) Позволяет создавать продукты пошагово.
2) Позволяет использовать один и тот же код для создания различных продуктов.
3) Изолирует сложный код сборки продукта от его основной бизнес-логики.
Минусы: 1) Усложняет код программы из-за введения дополнительных классов.
Пример использования: В игре пользователю необходимо выбрать готового персонажа из списка.
Каждому такому персонажу выставляются определенные параметры, 
недоступные для редактирования пользователю
*/
