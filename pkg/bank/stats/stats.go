package stats

import (
	"github.com/Azino-Beatcoin/bank/pkg/bank/types"
)

// Avg function is function
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for i := 0; i < len(payments); i++ {
		sum += payments[i].Amount
	}
	return sum / types.Money(len(payments))
}

// TotalInCategory function
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for i := 0; i < len(payments); i++ {
		if payments[i].Category == category {
			sum += payments[i].Amount
		}
	}
	return sum
}
