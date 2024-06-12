package pattern;

import (
	"fmt"
)

// State interface
type State interface {
    addItem(int) error
    requestItem() error
    dispenseItem() error
}

// NoItemState struct
type NoItemState struct {
    vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
    return fmt.Errorf("item out of stock")
}

func (i *NoItemState) addItem(count int) error {
    i.vendingMachine.incrementItemCount(count)
    i.vendingMachine.setState(i.vendingMachine.hasItem)
    return nil
}

func (i *NoItemState) dispenseItem() error {
    return fmt.Errorf("item out of stock")
}

// HasItemState struct
type HasItemState struct {
    vendingMachine *VendingMachine
}

func (i *HasItemState) requestItem() error {
    if i.vendingMachine.itemCount == 0 {
        i.vendingMachine.setState(i.vendingMachine.noItem)
        return fmt.Errorf("no item present")
    }
    fmt.Printf("Item requestd\n")
    i.vendingMachine.setState(i.vendingMachine.itemRequested)
    return nil
}

func (i *HasItemState) addItem(count int) error {
    fmt.Printf("%d items added\n", count)
    i.vendingMachine.incrementItemCount(count)
    return nil
}

func (i *HasItemState) dispenseItem() error {
    return fmt.Errorf("please select item first")
}

// ItemRequestedState struct
type ItemRequestedState struct {
    vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
    return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
    return fmt.Errorf("item Dispense in progress")
}

func (i *ItemRequestedState) dispenseItem() error {
    return fmt.Errorf("please insert money first")
}

// VendingMachine struct
type VendingMachine struct {
    hasItem       State
    itemRequested State
    noItem        State

    currentState State

    itemCount int
    itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
    v := &VendingMachine{
        itemCount: itemCount,
        itemPrice: itemPrice,
    }
    hasItemState := &HasItemState{
        vendingMachine: v,
    }
    itemRequestedState := &ItemRequestedState{
        vendingMachine: v,
    }
    noItemState := &NoItemState{
        vendingMachine: v,
    }

    v.setState(hasItemState)
    v.hasItem = hasItemState
    v.itemRequested = itemRequestedState
    v.noItem = noItemState
    return v
}

func (v *VendingMachine) requestItem() error {
    return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
    return v.currentState.addItem(count)
}

func (v *VendingMachine) dispenseItem() error {
    return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
    v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
    fmt.Printf("Adding %d items\n", count)
    v.itemCount = v.itemCount + count
}

// StateUserCode Пользовательский код использования паттерна состояние
func StateUserCode() {
    vendingMachine := newVendingMachine(1, 10)

    err := vendingMachine.requestItem()
    if err != nil {
        fmt.Println(err.Error())
    }

    err = vendingMachine.dispenseItem()
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Println()

    err = vendingMachine.addItem(2)
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Println()

    err = vendingMachine.requestItem()
    if err != nil {
        fmt.Println(err.Error())
    }

    err = vendingMachine.dispenseItem()
    if err != nil {
        fmt.Println(err.Error())
    }
}

/*
Паттерн состояние
Применимость: 1) Когда наш объект имеет много состояний и постоянно свапает их
2) Когда код класса содержит множество больших, похожих друг на друга, 
условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.
Плюсы: 1) Концентрирует в одном месте код состояния
2) Избавляет от множества больших условных операторов машины состояний.
Минусы: 1) Может усложнить код, если состояний мало и они редко меняются.
Пример использования: В любой РПГ игре существует множество различных механик.
Любой, даже незначительный npc может быть в нескольких состояниях, в зависимости от окружающего мира
Паттерн поможет справиться с введением новых состояний и изменение старых, не затрагивая рабочий код
*/