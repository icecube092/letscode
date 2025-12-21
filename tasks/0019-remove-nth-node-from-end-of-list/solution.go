/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var prev, next *ListNode
    var all, cur int
    var find func (node *ListNode)

    find = func (node *ListNode) {
        if node == nil {
            return 
        }

        all++
        cur++
        find(node.Next)
        cur--

        if all - cur == n-1 {
            next = node
        }
        if all - cur == n+1 {
            prev = node
        }
    }

    find(head)

    if prev == nil {
        if next != nil {
            return next
        }
        return nil
    }

    prev.Next = next

    return head

}
