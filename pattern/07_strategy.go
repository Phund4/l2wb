package pattern

import "fmt"

// EvictionAlgo interface
type EvictionAlgo interface {
	evict(c *Cache)
}

// FIFO struct
type FIFO struct {
}

// Метод вытеснения типа FIFO
func (*FIFO) evict(c *Cache) {
	fmt.Println("Evicting use FIFO")
}

// LRU struct
type LRU struct {
}

// Метод вытеснения типа LRU
func (*LRU) evict(c *Cache) {
	fmt.Println("Evicting use LRU")
}

// LFU struct
type LFU struct {

}

// Метод вытеснения типа LFU
func (*LFU) evict(c *Cache) {
	fmt.Println("Evicting use LFU");
}

// Cache struct
type Cache struct {
	storage map[string]string
	evictionAlgo EvictionAlgo
	capacity int
}

// NewCache Функция создания кэша
func NewCache(algo EvictionAlgo) *Cache {
	storage := make(map[string]string);
	return &Cache{
		storage: storage,
		evictionAlgo: algo,
		capacity: 0,
	}
}

// SetEvictionAlgo Метод изменения алгоритма
func (c *Cache) SetEvictionAlgo(algo EvictionAlgo) {
	c.evictionAlgo = algo
}

// Add Добавление в кэш
func (c *Cache) Add(key, value string) {
    c.Evict()
    c.capacity++
    c.storage[key] = value
}

// Get Получение ячейки кэша
func (c *Cache) Get(key string) {
    delete(c.storage, key)
}

// Evict Чистка кэша
func (c *Cache) Evict() {
    c.evictionAlgo.evict(c)
    c.capacity--
}

// StrategyUserCode Пользовательский код использования паттерна стратегия
func StrategyUserCode() {
	lfu := &LFU{}
    cache := NewCache(lfu)

    cache.Add("a", "1")
    cache.Add("b", "2")

    cache.Add("c", "3")

    lru := &LRU{}
    cache.SetEvictionAlgo(lru)

    cache.Add("d", "4")

    fifo := &FIFO{}
    cache.SetEvictionAlgo(fifo)

    cache.Add("e", "5")
}

/*
Паттерн стратегия
Применимость: 1) Когда нам нужно использовать несколько различных типов алгоритма внутри одного объекта
2) Когда мы не хотим показывать часть реализации другим абстракциям
3) Когда у нас есть множество похожих структур, отличающихся только некоторым поведением
Плюсы: 1) Замена алгоритмов на лету
2) Изолирует код и данные алгоритмов от других структур
3) Реализует принцип открытости/закрытости
Минусы: 1) Усложняет программу введением новых структур
2) Клиент должен знать принцип работы алгоритмов, чтобы выбрать подходящий
Пример использования: Выше приведен пример использования стратегии. Мы имеем несколько
алгоритмов очистки кэша, каждый из которых по своему хорош. Для динамичного ипользования 
разных типов применяется данный паттерн. Также присутствует удобство добавления нового алгоритма,
не затронув работу уже написанных.
*/