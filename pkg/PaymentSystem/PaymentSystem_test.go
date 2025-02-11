package PaymentSystem

import (
	"testing"
)

func TestEmition(t *testing.T) {
	ps := New()
	err := ps.Emition("BYN", 10000)

	if err != nil {
		t.Errorf("Emition error: %v", err)
	}

	emissionAccount := EmitionAccount("BYN")
	account, exists := ps.accounts[emissionAccount]

	if !exists {
		t.Errorf("Emition account does not exist")
	}

	if account.Balance != 10000 {
		t.Errorf("Expected 10000, get %f", account.Balance)
	}
}
