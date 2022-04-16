package palindrome

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) (bool, error) {
	lower := strings.ToLower(s)
	regex, err := regexp.Compile(`\W`)
	if err != nil {
		return false, err
	}
	bytes := regex.ReplaceAll([]byte(lower), []byte(""))
	lower = string(bytes)
	var reversed []string
	for _, l := range strings.Split(lower, "") {
		reversed = insertAtBeginning(l, reversed)
	}
	return strings.Join(reversed, "") == lower, nil
}

func insertAtBeginning(item string, slice []string) []string {
	newSlice := []string{item}
	return append(newSlice, slice...)
}
