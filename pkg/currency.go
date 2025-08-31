package pkg

import (
	"strings"
)

// MapStableCoinToISO converts stablecoin to fiat ISO code
// func MapStablecoinToISO(currency string) string {
// 	switch currency {
// 	case "cNGN", "cngn": return "ngn"
// 	case "cXAF", "cxaf": return "xaf"
// 	case "USDx", "usdx": return "usd"
// 	case "EURx", "eurx": return "eur"
// 	case "cZAR", "czar": return "zar"
// 	case "cGHS", "cghs": return "ghs"
// 	case "cKES", "kes": return "kes"
// 	default: return currency
// 	}
// }

var currency = []string{"ngn", "usd", "eur", "xaf"}

func MapStablecoinToISO(coin string) string {
	coinLower := strings.ToLower(coin)

	if len(coinLower) <= 3 {
		return coinLower
	}

	for i:=0; i < 4; i++{
		if strings.Contains(coinLower,currency[i]) {
			return currency[i]
		}
	}
	return coinLower
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