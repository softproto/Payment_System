package PaymentSystem

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

var EmissionAccounts = map[string]string{
	"BYN": "BY13NBRB3600900000002Z00AB00",
	"USD": "BY13NBRB3600900000002Z00AB01",
}

var WithdrawalAccounts = map[string]string{
	"BYN": "BY13NBRB3600900000002Z00AB10",
	"USD": "BY13NBRB3600900000002Z00AB11",
}

type Subject struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type Account struct {
	IBAN     string  `json:"iban"`
	Currency string  `json:"currency"`
	Owner    Subject `json:"owner"`
	Balance  float64 `json:"balance"`
	Status   string  `json:"status"`
}

type Payment struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

type PaymentSystem struct {
	accounts map[string]*Account
}

// CreateAccount() создает новый счет с указанным IBAN, именем владельца и статусом
func (ps *PaymentSystem) CreateAccount(iban, name, currency, status string) error {

	if _, exists := ps.accounts[iban]; exists {
		return fmt.Errorf("IBAN already exists")
	}

	newAccount := &Account{
		IBAN:     iban,
		Currency: currency,
		Owner: Subject{
			ID:   uuid.New(),
			Name: name,
		},
		Balance: 0,
		Status:  status,
	}
	ps.accounts[iban] = newAccount

	fmt.Printf("Account %s (%s) created\n", iban, name)
	return nil
}

// Transfer() перевод между счетами
func (ps *PaymentSystem) Transfer(sender, recipient string, amount float64) error {
	s, exists := ps.accounts[sender]
	if !exists {
		return fmt.Errorf("sender's account is not exist")
	}

	if s.Status != "active" {
		return fmt.Errorf("sender's account is not active")
	}

	r, exists := ps.accounts[recipient]
	if !exists {
		return fmt.Errorf("recipient's account is not exist")
	}

	if r.Status != "active" {
		return fmt.Errorf("recipient's account is not active")
	}

	if r.Currency != s.Currency {
		return fmt.Errorf("currency mismatch")
	}

	if s.Balance < amount {
		return fmt.Errorf("insufficient funds in the account %s", sender)
	}

	s.Balance -= amount
	r.Balance += amount

	fmt.Printf("Transfer for %.2f %s from %s (%s) to %s (%s) complited\n", amount, s.Currency, s.IBAN, s.Owner.Name, r.IBAN, r.Owner.Name)
	return nil
}

// Payment() выполняет перевод денег через JSON-структуру.
func (ps *PaymentSystem) Payment(data []byte) error {

	var transfer Payment
	err := json.Unmarshal(data, &transfer)
	if err != nil {
		return fmt.Errorf("parsing JSON error: %v", err)
	}

	return ps.Transfer(transfer.Sender, transfer.Recipient, transfer.Amount)
}

// Emition() эмиссия денежных средств в указанной валюте
func (ps *PaymentSystem) Emition(currency string, amount float64) error {

	account, exists := ps.accounts[EmissionAccounts[currency]]
	if !exists {
		return fmt.Errorf("%s emition account is not exist", currency)
	}

	if account.Status != "active" {
		return fmt.Errorf("emition account is not active")
	}

	account.Balance += amount

	fmt.Printf("Emition for %.2f %s complited\n", amount, currency)
	return nil
}

// Withdrowal() вывод денежных средств из оборота с указанного счета в валюте источника
func (ps *PaymentSystem) Withdrowal(sender string, amount float64) error {

	return ps.Transfer(sender, WithdrawalAccounts[ps.accounts[sender].Currency], amount)

}

// New()
// 1. Создает экземпляр объекта PaymentSystem
// 2. Создает счета для эмисии денежных средств (для всех доступных валют)
// 3. Создает счета для вывода денежных средств из обращения (для всех доступных валют)
func New() *PaymentSystem {

	ps := &PaymentSystem{
		accounts: make(map[string]*Account),
	}
	for c, a := range EmissionAccounts {
		ps.CreateAccount(a, "EMITION", c, "active")
	}

	for c, a := range WithdrawalAccounts {
		ps.CreateAccount(a, "WITHDROWAL", c, "active")
	}

	return ps
}
