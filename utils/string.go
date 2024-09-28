package utils

import "strings"

func HasHttpPrefix(input string) bool {
	return strings.HasPrefix(input, "http://")
}

func HasHttpsPrefix(input string) bool {
	return strings.HasPrefix(input, "https://")
}

func RemoveProtocolFromUrl(s string) string {
	httpPrefix := "http://"
	if strings.HasPrefix(s, httpPrefix) {
		sub := strings.Split(s, httpPrefix)
		return sub[1]
	}

	httpsPrefix := "https://"
	if strings.HasPrefix(s, httpsPrefix) {
		sub := strings.Split(s, httpsPrefix)
		return sub[1]
	}

	return s
}
