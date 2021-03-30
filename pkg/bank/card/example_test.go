package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN: "01",
			Active: true,
			Balance: 5_000_00,
		},
		{
			PAN: "02",
			Active: false,
			Balance: 5_000_00,
		},
		{
			PAN: "03",
			Active: true,
			Balance: 0,
		},
		{
			PAN: "04",
			Active: true,
			Balance: 9_000_00,
		},
	}
	payments := PaymentSources(cards)
	
	fmt.Println(payments)
	
	//Output: 
	//[{auto 01 500000} {auto 04 900000}]
	
}