package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Zekeriyyah/stellar-x/pkg"
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

type FXService struct{}

func NewFXService() *FXService {
	return &FXService{}
}

// GetRate returns live FX rate between two currencies
func (s *FXService) GetRate(from, to string) (float64, error) {
	// Map stablecoins to fiat
	fiatFrom := pkg.MapStablecoinToISO(from)
	fiatTo := pkg.MapStablecoinToISO(to)

	if fiatFrom == fiatTo {
		return 1.0, nil
	}

	// Prefer CoinGecko for African currencies
	if pkg.IsAfricanCurrency(fiatFrom) || pkg.IsAfricanCurrency(fiatTo) {
		return s.getCoinGeckoRate(fiatFrom, fiatTo)
	}

	// Use Frankfurter for global fiat pairs
	return s.getFrankfurterRate(fiatFrom, fiatTo)
}

// getFrankfurterRate fetches rate from api.frankfurter.dev
func (s *FXService) getFrankfurterRate(from, to string) (float64, error) {
	url := fmt.Sprintf("https://api.frankfurter.dev/v1/latest?base=%s&symbols=%s", from, to)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var fxResp FrankfurterResponse
	if err := json.Unmarshal(body, &fxResp); err != nil {
		return 0, err
	}

	if rate, exists := fxResp.Rates[to]; exists {
		return rate, nil
	}
	return 0, fmt.Errorf("rate not found: %s/%s", from, to)
}

// getCoinGeckoRate fetches rate from CoinGecko
func (s *FXService) getCoinGeckoRate(from, to string) (float64, error) {
	url := "https://api.coingecko.com/api/v3/exchange_rates"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var cgResp CoinGeckoResponse
	if err := json.Unmarshal(body, &cgResp); err != nil {
		return 0, err
	}

	// Get values
	fromVal, fromExists := cgResp.Rates[from]
	toVal, toExists := cgResp.Rates[to]

	if !fromExists {
		return 0, fmt.Errorf("currency not found: %s", from)
	}
	if !toExists {
		return 0, fmt.Errorf("currency not found: %s", to)
	}

	// Convert: (from / USD) / (to / USD) = from/to
	rate := fromVal.Value / toVal.Value
	return rate, nil
}
