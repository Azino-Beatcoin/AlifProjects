package types

// Money type int64
type Money int64

// Currency type string
type Currency string

// Currency constants
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN type string
type PAN string

// Card struct template
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

// Payment struct
type Payment struct {
	ID     int
	Amount Money
}

// PaymentSource struct
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
