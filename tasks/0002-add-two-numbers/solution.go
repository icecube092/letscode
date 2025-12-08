/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    cur := &ListNode{}
    root := cur
    n1 := l1
    n2 := l2

    for {
        var n1val, n2val int
        if n1 != nil {
            n1val = n1.Val
            n1 = n1.Next
        }
        if n2 != nil {
            n2val = n2.Val
            n2 = n2.Next
        }

        if cur == nil {
            break
        }

        val := n1val + n2val
        if val >= 10 {
            cur.Val = val - 10 + cur.Val
            cur.Next = &ListNode{Val: 1}
        } else {
            if val + cur.Val >= 10 {
                cur.Val = val - 10 + cur.Val
                cur.Next = &ListNode{Val: 1}
            } else {
                cur.Val = val + cur.Val
                if n1 != nil || n2 != nil {
                    cur.Next = &ListNode{}
                }
            }
        }

        cur = cur.Next
    }

    return root
}
