package leetcode

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := ListNode{Val: 2}
	l2 := ListNode{Val: 1}
	fmt.Println(mergeTwoLists(&l1, &l2))
}
