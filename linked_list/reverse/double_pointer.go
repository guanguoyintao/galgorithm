package reverse

import "github.com/guanguoyintao/galgorithm/linked_list"

func DoublePointerReverseList[T any](head *linked_list.ListNode[T]) *linked_list.ListNode[T] {
	if head == nil {
		return nil
	}

	var newHead *linked_list.ListNode[T]
	for head != nil {
		node := head.Next
		head.Next = newHead
		newHead = head
		head = node
	}

	return newHead
}
