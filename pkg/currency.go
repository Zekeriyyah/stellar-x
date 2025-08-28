package pkg

import (
	"strings"
)

// MapCurrencyToISO converts stablecoin to fiat ISO code
func MapStablecoinToISO(currency string) string {
	switch currency {
	case "cNGN": return "NGN"
	case "cXAF": return "XAF"
	case "USDx": return "USD"
	case "EURx": return "EUR"
	default: return currency
	}
}


// IsAfricanCurrency checks if currency is African
func IsAfricanCurrency(currency string) bool {
	return currency == "ngn" || currency == "xaf" || currency == "zar" || currency == "kes" || currency == "ghs"
}

func ScrapeQuery(query string) []string {
	currencies := []string{"NGN", "XAF", "USD", "EUR", "ngn", "xaf", "usd", "eur"}
	
	presentCurr := []string{}
	for i:=0; i < len(currencies); i++ {
		if strings.Contains(query, currencies[i]) {
			presentCurr = append(presentCurr, currencies[i])
		}
	}
	return presentCurr
}