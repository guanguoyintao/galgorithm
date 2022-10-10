package reverse

import "github.com/guanguoyintao/galgorithm/linked_list"

func RecursionReverseList[T any](head *linked_list.ListNode[T]) *linked_list.ListNode[T] {
	if head == nil || head.Next == nil {
		return head
	}

	ret := RecursionReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return ret
}
