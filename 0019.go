func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    l := dummy
    r := dummy
    for i:=0; i < n; i ++ {
        r = r.Next
    }
    for r.Next != nil {
        r = r.Next
        l = l.Next
    }
    l.Next = l.Next.Next
    return dummy.Next
}