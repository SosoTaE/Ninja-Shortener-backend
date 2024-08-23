package main

import (
	"math/rand"
	"net/url"
	"strings"
)

func random(min, max int) int {
	return min + rand.Intn(max-min)
}

//func isURLValid(urlToCheck string) bool {
//	// Regex pattern for URL validation.
//	// It's a simplified version, you can find more comprehensive ones online.
//	regexPattern := `^(https?://)?([a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\.)+[a-zA-Z]{2,}$`
//
//	isValid, _ := regexp.MatchString(regexPattern, urlToCheck)
//	return isValid
//}

func isURLValid(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true

}

func adjustHTTPS(rawURL string) string {
	if strings.HasPrefix(rawURL, "http://www.") {
		return rawURL
	}

	if strings.HasPrefix(rawURL, "www.") {
		newText := strings.TrimPrefix(rawURL, "www.")
		return "http://www." + newText
	}

	if strings.HasPrefix(rawURL, "https://www.") {
		newText := strings.TrimPrefix(rawURL, "https://www.")
		return "http://www." + newText
	}

	if strings.HasPrefix(rawURL, "https://") {
		newText := strings.TrimPrefix(rawURL, "https://")
		return "http://www." + newText
	}

	if strings.HasPrefix(rawURL, "http://") {
		newText := strings.TrimPrefix(rawURL, "http://")
		return "http://www." + newText
	}

	if !strings.HasPrefix(rawURL, "www.") || !strings.HasPrefix(rawURL, "https://www.") || !strings.HasPrefix(rawURL, "http://www.") {
		return "http://www." + rawURL
	}

	return rawURL
}
