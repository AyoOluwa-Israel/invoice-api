package utils

import "strings"


func ConvertEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

