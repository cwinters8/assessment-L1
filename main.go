package main

import (
	"fmt"
	"log"

	"assessment/palindrome"
)

func main() {
	for _, s := range []string{"bah", "Baab", "racecar", "race car!"} {
		check, err := palindrome.IsPalindrome(s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("is %s a palindrome? %v\n", s, check)
	}

}
