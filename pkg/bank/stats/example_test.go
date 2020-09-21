package stats

import (
	"fmt"

	"github.com/Azino-Beatcoin/bank/pkg/bank/types"
)

func ExampleAvg() {
	var payments = []types.Payment{
		types.Payment{
			ID:     1,
			Amount: 100,
		},
		types.Payment{
			ID:     2,
			Amount: 200,
		},
		types.Payment{
			ID:     3,
			Amount: 300,
		},
		types.Payment{
			ID:     4,
			Amount: 400,
		},
	}

	result := Avg(payments)

	fmt.Println(result)

	// Output:
	// 250
}
