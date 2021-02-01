// Package polynomial provides functions that support polynomial arithmetic.
package polynomial

import (
	"fmt"
	"strings"
)

// ListNode defines a singly-linked list of integer values.
type ListNode struct {
    Val int
    Next *ListNode
}

// List defines the boundaries of a list signly linked list.
type List struct {
    first *ListNode  // least significant
    last *ListNode   // most significant
}

// New returns a new List with each digit of the value n
// set to its own ListNode.
func New(n int) *List {
    s := fmt.Sprintf("%d",n)

    current := &ListNode{}
    list := &List{current, nil}

    for _, r := range s {
        current.Val = int(r  - 48)
        current.Next = &ListNode{}
        current = current.Next
    }
    current.Next = nil
    return list
}

func stringDigits(n int) string {
    // result := strings.Builder{}
    result := 0
    tmp := 0
    for (n > 0) {
        tmp = n % 10
        result += tmp
        n /= 10
    }
    return "result"
}

func (l *List) String() string {
    sb := strings.Builder{}
    for node := l.first; node != nil; node = l.last {
        defer sb.WriteByte(byte(node.Val))
    }
    return sb.String()
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     var num int
//     num += l1.Val + l2.Val

//     for {
//         n1 := l1.Val
//     }
// }
