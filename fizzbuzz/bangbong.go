package fizzbuzz

import (
	"fmt"
)

func FizzBuzzBangBong(n int) {
	for i := 1; i <= n; i++ {
		if check(i, 3) && check(i, 5) && check(i, 19) {
			print(i, "BONG")
		} else if check(i, 5) && check(i, 19) {
			print(i, "BUZZBANG")
		} else if check(i, 3) && check(i, 5) {
			print(i, "FIZZBUZZ")
		} else if check(i, 19) {
			print(i, "BANG")
		} else if check(i, 5) {
			print(i, "BUZZ")
		} else if check(i, 3) {
			print(i, "FIZZ")
		}
	}
}

func check(n int, multiple int) bool {
	return n%multiple == 0
}

func print(n int, s string) {
	fmt.Printf("%d: %s\n", n, s)
}
