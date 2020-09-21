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
