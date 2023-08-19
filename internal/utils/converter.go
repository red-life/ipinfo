package utils

import "strings"

func StringToBool(s string) bool {
	return strings.ToLower(s) == "true" || strings.ToLower(s) == "1"
}
