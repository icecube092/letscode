/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode]struct{})

    for {
        if _, ok := m[headA]; ok {
            return headA
        }
        if _, ok := m[headB]; ok {
            return headB
        }

        if headA != nil {
            m[headA] = struct{}{}
            if headA.Next != nil {
                headA = headA.Next
            } else {
                headA = nil
            }
        }

        if headB != nil {
            if _, ok := m[headB]; ok {
                return headB
            }
            m[headB] = struct{}{}
            if headB.Next != nil {
                headB = headB.Next
            } else {
                headB = nil
            }
        }

        if headA == nil && headB == nil {
            return nil
        }
    }

    return nil
}
