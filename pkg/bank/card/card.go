package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	
	number := []types.PaymentSource{}

	for _, card := range cards {
	
		if card.Balance > 0 && card.Active {

			number = append(number, types.PaymentSource{Type: "auto", Number: card.PAN, Balance: card.Balance}) 
		}	
	}
	return number
}