package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output:
	// 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 5_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)

	// Output:
	// 500000
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 10_000_00, Active: false}, 5_000_00)
	fmt.Println(result.Balance)

	// Output:
	// 1000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 10_000_00, Active: true}, 50_000_00)
	fmt.Println(result.Balance)

	// Output:
	// 1000000
}

func ExampleDeposit_positive() {
	card := types.Card{
		Balance: 10_000_00,
		Active:  true,
	}
	Deposit(&card, 20_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 3000000
}

func ExampleDeposit_inactive() {
	card := types.Card{
		Balance: 10_000_00,
		Active:  false,
	}
	Deposit(&card, 20_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 1000000
}

func ExampleDeposit_limit() {
	card := types.Card{
		Balance: 10_000_00,
		Active:  true,
	}
	Deposit(&card, 80_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 1000000
}
