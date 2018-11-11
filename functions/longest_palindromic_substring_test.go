package functions

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	resp := longestPalindromeV1("ffffdabad")
	fmt.Println(resp)
}
