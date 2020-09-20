package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		types.Card{
			Name:    "5555 5555 5555 5555",
			Balance: 100,
			Active:  true,
		},
		types.Card{
			Name:    "5555 5555 5555 5556",
			Balance: 100,
			Active:  false,
		},
		types.Card{
			Name:    "5555 5555 5555 5557",
			Balance: -100,
			Active:  true,
		},
		types.Card{
			Name:    "5555 5555 5555 5558",
			Balance: 200,
			Active:  true,
		},
	}
	result := PaymentSources(cards)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i].Number)
	}

	// Output:
	// 5555 5555 5555 5555
	// 5555 5555 5555 5558
}
