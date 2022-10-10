package reverse

import "github.com/guanguoyintao/galgorithm/linked_list"

func reverseList[T any](head *linked_list.ListNode[T]) *linked_list.ListNode[T] {
	if head == nil {
		return nil
	}

	cur := head
	for head.Next != nil {
		t := head.Next.Next
		head.Next.Next = cur // 反转原指针方向
		cur = head.Next      // 将新头节点移到下一位
		head.Next = t        // 连接回断开的地方，继续重复上面操作
	}

	return cur
}
