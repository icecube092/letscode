/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    lessHead := &ListNode{}
    lh := &ListNode{}
    greatHead := &ListNode{}
    gh := &ListNode{}

    for root := head; root != nil; root = root.Next {
        if root.Val < x {
            if lessHead.Next == nil {
                lh = root
                lessHead.Next = lh
            } else {
                lh.Next = root
                lh = root
            }
        } else {
            if greatHead.Next == nil {
                gh = root
                greatHead.Next = gh
            } else {
                gh.Next = root
                gh = root
            }
        }
    }

    gh.Next = nil

    lh.Next = greatHead.Next

    if lessHead.Next == nil {
        return greatHead.Next
    }

    return lessHead.Next
}
