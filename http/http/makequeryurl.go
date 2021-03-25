package http

import "strings"

// MakeQueryURL - join args into a google-style query string url
func MakeGoogleQueryURL(args []string, SearchStringPrefix string, SearchStringSep string) string {
	// https://www.google.com/search?q=trump+biden
	return SearchStringPrefix + strings.Join(args, SearchStringSep)
}
