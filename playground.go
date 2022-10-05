package main

import "fmt"

func isPalindrome(x int) bool {
	// When x < 0, x is not a palindrome
	// Also if the last digit of the number is 0, in order to be a palindrome,
	// the first digit of the number also needs to be 0.
	// Only 0 satisfy this property.

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	y := 0
	for x >= 1 {
		y = y*10 + x%10
		x = x / 10
	}
	fmt.Println(x)
	if y == x {
		return true
	} else {
		return false
	}
}

func main() {
	x := 1221
	fmt.Println(isPalindrome(x))
	fmt.Println(x)
}
