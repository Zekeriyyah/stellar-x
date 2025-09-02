package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// FXResponse for Frankfurter
type FrankfurterResponse struct {
	Base   string             `json:"base"`
	Rates  map[string]float64 `json:"rates"`
}

// CoinGeckoResponse for fiat rates
type CoinGeckoResponse struct {
	Rates map[string]struct {
		Name  string  `json:"name"`
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
		Type  string  `json:"type"`
	} `json:"rates"`
}

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



func GetAllFxCurrencies() (*CoinGeckoResponse, *FrankfurterResponse) {
	url := "https://api.coingecko.com/api/v3/exchange_rates"
	resp, err := http.Get(url)
	if err != nil {
		return &CoinGeckoResponse{}, &FrankfurterResponse{}
	}
	defer resp.Body.Close()

	cbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &CoinGeckoResponse{}, &FrankfurterResponse{}
	}

	cgResp := CoinGeckoResponse{}
	if err := json.Unmarshal(cbody, &cgResp); err != nil {
		return &CoinGeckoResponse{}, &FrankfurterResponse{}
	}

	url2 := "https://api.frankfurter.dev/v1/latest"
	resp2, err := http.Get(url2)
	if err != nil {
		return &CoinGeckoResponse{}, &FrankfurterResponse{}
	}
	defer resp2.Body.Close()

	fbody, err := io.ReadAll(resp2.Body)
	if err != nil {
		return &CoinGeckoResponse{}, &FrankfurterResponse{}
	}

	fxResp := FrankfurterResponse{}
	if err := json.Unmarshal(fbody, &fxResp); err != nil {
		return &CoinGeckoResponse{}, &FrankfurterResponse{}
	}




	return &cgResp, &fxResp
}