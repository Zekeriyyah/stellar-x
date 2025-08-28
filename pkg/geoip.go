package pkg

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetCountryFromIP fetches country name from IP using ipapi.co
// Returns "unknown" if request fails
func GetCountryFromIP(ip string) string {
	url := fmt.Sprintf("https://ipapi.co/%s/country_name/", ip)

	resp, err := http.Get(url)
	if err != nil {
		return "unknown"
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "unknown"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "unknown"
	}

	country := string(body)
	if country == "" {
		return "unknown"
	}

	return strings.TrimSpace(country)
}

