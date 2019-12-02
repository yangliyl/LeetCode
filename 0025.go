/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    prev, start, end, next := dummy, head, head, head

    for next != nil {
        for i := 1; i < k && end != nil; i ++ {
            end = end.Next
        }
        
        if end == nil {
            break
        }

        next = end.Next
        end.Next = nil
        end = start
        start = reverse(start)
        end.Next = next
        prev.Next = start
        prev, start, end = end, next, next
    }

    return dummy.Next
}

func reverse(head *ListNode) *ListNode {
    var prev, next *ListNode
    cur := head
    for cur != nil {
        next = cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
}