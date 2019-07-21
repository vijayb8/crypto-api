package web

import "github.com/microcosm-cc/bluemonday"

// Sanitize removes html tags except <img> and <a>
func Sanitize(html string) string {
	return bluemonday.StrictPolicy().Sanitize(html)
}
