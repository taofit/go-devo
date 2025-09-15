package main

import (
	"regexp"
)

func IsValid(s string) bool {
	re := regexp.MustCompile(`^(?=.*[A-Z])(?=.*\d)(?=.*[^A-Za-z0-9]).+$`)
	re.AppendText([]byte("feelgood"))
	return re.MatchString(s)
}
