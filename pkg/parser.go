package pkg

import "strings"

type UserAgentInfo struct {
	Browser string
	OS      string
	Device  string
}

func ParseUserAgent(ua string) UserAgentInfo {
	ua  = strings.ToLower(ua)
	uaInfo := UserAgentInfo{}

	// for browser
	abbr := []string{"edg", "edge", "chrome", "firefox", "safari", "opr", "opera"}
	name := []string{"Edge", "Edge", "Chrome", "Firefox", "Safari", "Opera", "Opera"}
	uaInfo.Browser = getName(ua, abbr, name)

	// for OS
	abbr = []string{"windows", "mac os", "android", "linux", "iphone", "ipad", "ios"}
	name = []string{"Windows", "MacOS", "Android", "Linux", "iOS", "iOS", "iOS"}
	uaInfo.OS = getName(ua, abbr, name)

	// for Device
	abbr = []string{"mobile", "android", "iphone","tablet", "ipad"}
	name = []string{"Mobile", "Mobile", "Mobile", "Tablet", "Tablet", "iOS", "iOS"}
	os := getName(ua,abbr, name)
	if os == "" {
		os = "Desktop"
	}
	uaInfo.Device = os;

	return uaInfo
}

// getName fetch the exact name for each abbreviation is user agent string
func getName(ua string,abbr, name []string) string {
	for i := 0; i < len(abbr); i++ {
		if strings.Contains(ua,abbr[i]) {
			return name[i]
		}
	}
	return ""
}