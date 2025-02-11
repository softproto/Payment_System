package PaymentSystem

import (
	"fmt"

	"github.com/google/uuid"
)

const emissionAccount = "BY13NBRB3600900000002Z00AB00"
const withdrawalAccount = "BY13NBRB3600900000002Z00AB01"

type Subject struct {
	ID   uuid.UUID
	Name string
}

type Account struct {
	IBAN    string  `json:"iban"`
	Owner   Subject `json:"owner"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
}

type PaymentSystem struct {
	accounts map[string]*Account
}

// CreateAccount() создает новый счет с указанным IBAN, именем владельца и статусом
func (ps *PaymentSystem) CreateAccount(iban, name, status string) (*Account, error) {

	if _, exists := ps.accounts[iban]; exists {
		return nil, fmt.Errorf("IBAN exists")
	}

	newAccount := &Account{
		IBAN: iban,
		Owner: Subject{
			ID:   uuid.New(),
			Name: name,
		},
		Balance: 0,
		Status:  status,
	}
	ps.accounts[iban] = newAccount
	return newAccount, nil
}

func (ps *PaymentSystem) emition(amount float64) error {

	account := ps.accounts[emissionAccount]
	if account.Status != "active" {
		return fmt.Errorf("account is not active!")
	}

	account.Balance += amount
	return nil
}

func New(amount float64) *PaymentSystem {
	ps := &PaymentSystem{
		accounts: make(map[string]*Account),
	}

	ps.CreateAccount(emissionAccount, "NBRB", "active")
	ps.emition(amount)
	// ps.accounts[emission] = &Account{
	// 	IBAN: emission,
	// 	Owner: Subject{
	// 		ID:   uuid.New(),
	// 		Name: "NBRB",
	// 	},
	// 	Balance: amount,
	// 	Status:  "active",
	// }
	ps.CreateAccount(withdrawalAccount, "NBRB", "active")
	// ps.accounts[withdrawal] = &Account{
	// 	IBAN: withdrawal,
	// 	Owner: Subject{
	// 		ID:   uuid.New(),
	// 		Name: "NBRB",
	// 	},
	// 	Balance: 0,
	// 	Status:  "active",
	// }

	return ps
}
