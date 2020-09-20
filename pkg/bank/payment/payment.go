package payment

import (
	"bank/pkg/bank/types"
)

// Max returns maximal payment
func Max(payments []types.Payment) types.Payment {
	max := payments[0]

	for i := 0; i < len(payments); i++ {
		if payments[i].Amount > max.Amount {
			max = payments[i]
		}
	}

	return max
}

// PaymentSources function
func PaymentSources(cards []types.Card) []types.PaymentSource {
	arr := make([]types.PaymentSource, 0)
	for i := 0; i < len(cards); i++ {
		if cards[i].Balance > 0 && cards[i].Active == true {
			obj := types.PaymentSource{
				Type:    "card",
				Number:  cards[i].Name,
				Balance: cards[i].Balance,
			}
			arr = append(arr, obj)
		}
	}
	return arr
}
