package main

import "testing"

func TestPalindromeNumber(t *testing.T) {
	result := isPalindrome(121)

	if result != true {
		t.Error("Result is not as expected")
	}
}
