package linked_list

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func NewListNode[T any](list []T) *ListNode[T] {
	node := &ListNode[T]{}
	l := &ListNode[T]{Next: node}
	for _, v := range list {
		node.Val = v
		node.Next = &ListNode[T]{}
		node = node.Next
	}

	return l
}
