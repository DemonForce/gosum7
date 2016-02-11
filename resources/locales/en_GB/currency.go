package en_GB

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "ILR", DisplayName: "Israeli Shekel (1980–1985)", Symbol: ""},
		{Currency: "ILS", DisplayName: "Israeli New Shekel", Symbol: ""},
	}
}
