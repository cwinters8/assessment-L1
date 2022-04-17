package palindrome

import (
	"math"
	"regexp"
	"strings"
)

func IsPalindrome(s string) (bool, error) {
	regex, err := regexp.Compile(`\W`)
	if err != nil {
		return false, err
	}
	length := len(s)
	halfLength := int(math.Round(float64(length) / 2))
	compareIdx := 0
	for i := length - 1; i >= halfLength; i-- {
		if regex.Match([]byte{s[i]}) || regex.Match([]byte{s[compareIdx]}) {
			continue
		}
		currentLetter := strings.ToLower(string(s[i]))
		compareLetter := strings.ToLower(string(s[compareIdx]))

		if currentLetter != compareLetter {
			return false, nil
		}
		compareIdx++
	}
	return true, nil
}
