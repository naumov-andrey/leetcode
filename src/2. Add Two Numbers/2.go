/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    var res *ListNode
    p1, p2, t := l1, l2, res
    
    for p1 != nil && p2 != nil {
        if t == nil {
            t = &ListNode{}
            res = t
        } else {
            t.Next = &ListNode{}
            t = t.Next
        }
        
        t.Val = p1.Val + p2.Val + carry
        carry = t.Val / 10
        if t.Val >= 10 {
            t.Val %= 10
        }
        
        p1, p2 = p1.Next, p2.Next
    }
    
    var rest *ListNode
    if p1 != nil {
        rest = p1
    } else if p2 != nil {
        rest = p2
    }
    
    for rest != nil {
        t.Next = &ListNode{}
        t = t.Next
        
        t.Val = rest.Val + carry
        carry = t.Val / 10
        if t.Val >= 10 {
            t.Val %= 10
        }
        
        rest = rest.Next
    }
    
    if carry != 0 {
        t.Next = &ListNode{carry, nil}
    }
    
    return res
}
