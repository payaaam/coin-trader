package utils

import (
	"strings"
)

func Normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}
