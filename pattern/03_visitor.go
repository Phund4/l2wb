package pattern;

import (
	"fmt"
)

// Visitor interface
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// Shape interface
type Shape interface {
	getType() string
	accept(Visitor)
}

// Square struct
type Square struct {
	side int
}

func (s *Square) getType() string {
	return "square";
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

// Circle struct
type Circle struct {
	radius int
}

func (c *Circle) getType() string {
	return "circle";
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c);
}

// Rectangle struct
type Rectangle struct {
	length int
	width int
}

func (r *Rectangle) getType() string {
	return "rectangle";
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r);
}

// AreaCalculator struct
type AreaCalculator struct {
	area int;
}

// Метод подсчета площади квадрата
func (ac *AreaCalculator) visitForSquare(s *Square) {
	ac.area = s.side * s.side
	fmt.Printf("Square area: %d\n", ac.area)
}

// Метод подсчета площади круга
func (ac *AreaCalculator) visitForCircle(c *Circle) {
	ac.area = int(3.14 * float64(c.radius) * float64(c.radius))
	fmt.Printf("Circle area: %d\n", ac.area)
}

// Метод подсчета площади прямоугольника
func (ac *AreaCalculator) visitForRectangle(r *Rectangle) {
	ac.area = r.length * r.width;
	fmt.Printf("Rectangle area: %d\n", ac.area)
}

// VisitorUserCode Пользовательский код использования паттерна посетитель
func VisitorUserCode() {
	square := &Square{side: 2}
    circle := &Circle{radius: 3}
    rectangle := &Rectangle{length: 2, width: 3}

    areaCalculator := &AreaCalculator{}
    square.accept(areaCalculator)
    circle.accept(areaCalculator)
    rectangle.accept(areaCalculator)
}

/*
Паттерн посетитель
Применимость: 1) Нам необходимо выполнить какое то действие над всеми объектами сложной структуры
2) Нововведение имеет смысл только для некоторых объектов в иерархии
Плюсы: 1) Упрощает добавление операций, работающих со сложными структурами объектов.
2) Объединяет родственные операции в одном классе.
3) Посетитель может накапливать состояние при обходе структуры элементов.
Минусы: 1) Паттерн не применим, если иерархия элементов часто меняется. 
Пример использования: Важно использовать такой паттерн в случае, когда иерархия элементов меняется
нечасто, зато поведение для объектов этой иерархии нужно менять постоянно.
Примером является интернет магазин с различными товарами, и нам нужно с каждым обновлением
вносить новые метрики качества товара.
*/