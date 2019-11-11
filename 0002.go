/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    p,q := l1,l2
    up := 0
    tmp := dummyHead
    for p != nil || q != nil || up > 0 {
        x, y := 0, 0
        if p != nil {
            x = p.Val
            p = p.Next
        }
        if q != nil {
            y = q.Val
            q = q.Next
        }
        sum := x + y + up
        up = sum / 10
        tmp.Next = &ListNode{Val: sum % 10}
        tmp = tmp.Next
    }
    return dummyHead.Next
}