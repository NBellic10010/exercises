package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Insert(pre *ListNode, p *ListNode) {
	var iter *ListNode
	iter = pre.Next
	for iter != nil {
		if iter.Val >= p.Val && pre.Val <= p.Val {
			pre.Next = p
			p.Next = iter
		} else {
			pre = iter
			iter = iter.Next
		}
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	var p *ListNode
	var pre *ListNode
	var next *ListNode
	var head *ListNode
	head = lists[0]
	for li := 1; li < len(lists); li++ {
		p = lists[li]
		pre = lists[li-1]
		for p != nil {
			next = p.Next
			Insert(pre, p)
			p = next
		}
	}
	return head
}
