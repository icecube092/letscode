/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    var nextF, nextS = head, head

    for {
        if nextF.Next == nil {
            return false
        }

        nextF = nextF.Next

        if nextF.Next == nil {
            return false
        }

        nextF = nextF.Next
        nextS = nextS.Next

        if nextS == nextF {
            return true
        }
    }

    return false
}
