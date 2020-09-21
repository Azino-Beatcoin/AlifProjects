package card

import (
	"github.com/Azino-Beatcoin/bank/pkg/bank/types"
)

// IssueCard returns card type from types.Card template
func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}

// Withdraw collects amount from card balance
func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Balance <= 0 {
		return card
	}

	if amount <= 0 {
		return card
	}

	if card.Active == false {
		return card
	}

	if card.Balance-amount < 0 {
		return card
	}

	if amount > 20_000_00 {
		return card
	}

	card.Balance = card.Balance - amount

	return card
}

// Deposit function add money to our wallet
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}

	if amount > 5000000 || amount < 0 {
		return
	}

	card.Balance += amount
}

// Total func
func Total(cards []types.Card) types.Money {
	var sum types.Money = types.Money(0)

	if len(cards) == 0 {
		return 0
	}

	for i := 0; i < len(cards); i++ {
		sum += cards[i].Balance
	}

	return sum
}

// PaymentSources function
func PaymentSources(cards []types.Card) []types.PaymentSource {
	arr := make([]types.PaymentSource, 0)
	for i := 0; i < len(cards); i++ {
		if cards[i].Balance > 0 && cards[i].Active == true {
			obj := types.PaymentSource{
				Type:    "card",
				Number:  string(cards[i].PAN),
				Balance: cards[i].Balance,
			}
			arr = append(arr, obj)
		}
	}
	return arr
}