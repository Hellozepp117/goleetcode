package functions_test

import (
	"fmt"
	"leetcode/functions"
	"testing"
)

func TestAddTwoNums(t *testing.T) {
	l1 := &functions.ListNode{Val: 5}
	// l2 := &functions.ListNode{Val: 0, Next: l1}
	// l3 := &functions.ListNode{Val: 0, Next: l2}
	// l4 := &functions.ListNode{Val: 0, Next: l3}
	// l5 := &functions.ListNode{Val: 0, Next: l4}
	// l6 := &functions.ListNode{Val: 0, Next: l5}
	// l7 := &functions.ListNode{Val: 0, Next: l6}
	// l8 := &functions.ListNode{Val: 0, Next: l7}
	// l9 := &functions.ListNode{Val: 0, Next: l8}
	// l10 := &functions.ListNode{Val: 0, Next: l9}
	// l11 := &functions.ListNode{Val: 0, Next: l10}
	// l12 := &functions.ListNode{Val: 0, Next: l11}
	// l13 := &functions.ListNode{Val: 0, Next: l12}
	// l14 := &functions.ListNode{Val: 0, Next: l13}
	// l15 := &functions.ListNode{Val: 0, Next: l14}
	// l16 := &functions.ListNode{Val: 0, Next: l15}
	// l17 := &functions.ListNode{Val: 0, Next: l16}
	// l18 := &functions.ListNode{Val: 0, Next: l17}
	// l19 := &functions.ListNode{Val: 0, Next: l18}
	// l20 := &functions.ListNode{Val: 0, Next: l19}
	// l21 := &functions.ListNode{Val: 0, Next: l20}
	// l22 := &functions.ListNode{Val: 0, Next: l21}
	// l23 := &functions.ListNode{Val: 0, Next: l22}
	// l24 := &functions.ListNode{Val: 0, Next: l23}
	// l25 := &functions.ListNode{Val: 0, Next: l24}
	// l26 := &functions.ListNode{Val: 0, Next: l25}
	// l27 := &functions.ListNode{Val: 0, Next: l26}
	// l28 := &functions.ListNode{Val: 0, Next: l27}
	// l29 := &functions.ListNode{Val: 0, Next: l28}
	// l30 := &functions.ListNode{Val: 0, Next: l29}
	// l31 := &functions.ListNode{Val: 1, Next: l30}
	n1 := &functions.ListNode{Val: 5}
	// n2 := &functions.ListNode{Val: 6, Next: n1}
	// n3 := &functions.ListNode{Val: 4, Next: n2}
	result := functions.AddTwoNumbersV2(l1, n1)
	fmt.Println(result)
}
