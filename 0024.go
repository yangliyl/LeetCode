/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    dummy := &ListNode{}
    dummy.Next = head
    cur := dummy

    for cur.Next != nil && cur.Next.Next != nil {
        prev := cur.Next
        next := cur.Next.Next
        cur.Next = next
        prev.Next = next.Next
        next.Next = prev
        cur = cur.Next.Next
    }

    return dummy.Next
}