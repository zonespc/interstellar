package number

import (
	"fmt"
	"testing"
)

func TestReverseDigits(t *testing.T) {
	m := -102
	fmt.Println(ReverseDigits(m))
}

func TestAtoi(t *testing.T) {
	fmt.Println(Atoi("-91283472332"))
}

func TestIsPalindrome(t *testing.T) {
	x := -122
	fmt.Println(IsPalindrome(x))
}

func TestIsPalindromeByStr(t *testing.T) {
	x := 121
	fmt.Println(IsPalindromeByStr(x))
}

func TestIsMatchByRecall(t *testing.T) {
	s := "ab"
	p := ".*"
	fmt.Println(IsMatchByRecall(s, p))
}

func TestIsMatchByDP(t *testing.T) {
	s := "mississippi"
	p := "mis*is*ip*."
	fmt.Println(IsMatchByDP(s, p))
}
