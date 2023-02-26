/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    for root := head; root.Next != nil; {
        if root.Next.Val == root.Val {
            root.Next = root.Next.Next
            continue
        }

        root = root.Next
    }

    return head
}
