package types

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)


type Card struct {
	ID        int
	PAN       string
	Balance   Money
	Currency  Currency
	Color     string
	Name      string
	Active    bool
	MinBalance Money
}

type PaymentSource struct {
	Type     string
	Number   string
	Balance  Money
}