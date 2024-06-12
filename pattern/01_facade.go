package pattern;

import (
	"fmt"
)

// WalletFacade struct
type WalletFacade struct {
	account *Account
	wallet *Wallet
	securityCode *SecurityCode
}

// NewWalletFacade Функция создания нового фасада
func NewWalletFacade(accountName string, balance, accountCode int) *WalletFacade {
	fmt.Println("Create facade")
	return &WalletFacade{
		account: NewAccount(accountName),
		wallet: NewWallet(balance),
		securityCode: NewSecurityCode(accountCode),
	}
}

// Метод пополнения кошелька
func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting add money to wallet")

    err := w.account.checkAccount(accountID)
    if err != nil {
        return err
    }

    err = w.securityCode.checkCode(securityCode)
    if err != nil {
        return err
    }

    w.wallet.creditBalance(amount)
    return nil
}

// Метод снятие денег с кошелька
func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting debit money from wallet")

    err := w.account.checkAccount(accountID)
    if err != nil {
        return err
    }

    err = w.securityCode.checkCode(securityCode)
    if err != nil {
        return err
    }

    err = w.wallet.debitBalance(amount)
    if err != nil {
        return err
    }

    return nil
}


// Account struct
type Account struct {
	Name string
}

// NewAccount Функция создания аккаунта
func NewAccount(name string) *Account {
	fmt.Println("Create account")
	return &Account{
		Name: name,
	}
}

// Метод проверки входящего имени 
func (a *Account) checkAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("incorrect account name")
	}
	return nil;
}

// Wallet struct
type Wallet struct {
	Balance int
}

// NewWallet Функция создания кошелька
func NewWallet(balance int) *Wallet {
	fmt.Println("Create wallet")
	return &Wallet{
		Balance: balance,
	}
}

// Метод пополнения кошелька
func (w *Wallet) creditBalance(amount int) {
	w.Balance += amount;
}

// Метод снятия денег с кошелька
func (w *Wallet) debitBalance(amount int) error {
	if w.Balance < amount {
		return fmt.Errorf("Wallet balance isn't sufficient");
	}
	w.Balance -= amount;
	return nil;
}

// SecurityCode struct
type SecurityCode struct {
	Code int
}

// NewSecurityCode Функция создания кода безопасности
func NewSecurityCode(code int) *SecurityCode {
	fmt.Println("Create security code")
	return &SecurityCode{
		Code: code,
	}
}

// Метод проверки входящего кода безопасности
func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.Code != incomingCode {
		return fmt.Errorf("incorrect incoming code");
	}
	return nil;
}

// FacadeUserCode Пользовательский код использования паттерна фасад
func FacadeUserCode() {
    walletFacade := NewWalletFacade("abc", 1234, 20)

    err := walletFacade.addMoneyToWallet("abc", 1234, 10)
    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
    }

    err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
    }
}

/*
Паттерн фасад
Применимость: 1) В случае, когда в программе большая и сложная подсистема,
реализацию которой необходимо скрыть по причине ее сложности для пользователя.
Фасад предоставляет юзеру простой и понятный UX, скрывающий сложные бизнес процессы.
2) В случае, когда мы хотим разделить слои программы и заставить их общаться через
некую сущность, этой сущностью будет фасад.
Плюсы: 1) Изолирует клиентов от компонентов сложной подсистемы.
Минусы: 1) Фасад рискует быть привязанным ко всем структурам программы.
Пример использования: Выше в коде представлен пример использования фасада.
Пользователю не обязательно знать, как устроены бизнес процессы, протекающие при 
взаимодействии с сервисом оплаты. Мы предоставляем только два метода пополнения 
и снятия денег с его счета, оставляя реализация всех проверок под капотом.
*/