// 21. 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。
// 可以定义一个函数，接收两个链表的头节点作为参数，在函数内部使用双指针法，通过比较两个链表节点的值，将较小值的节点添加到新链表中，直到其中一个链表为空，然后将另一个链表剩余的节点添加到新链表中。
package task

type ListNode struct {
	Val  int
	Next *ListNode
}

func UnitLinklist(L1, L2 *ListNode) *ListNode {
	// if L1==nil{return L2}
	// if L2==nil{return L1}
	var L0, r1, r2, s *ListNode
	L0 = &ListNode{Val: -1}
	r1 = L1
	r2 = L2
	// if r1.val<r2.val {
	//     L0=r1
	//     r1=r1.next
	// }else{
	//     L0=r2
	//     r2=r2.next
	// }
	s = L0
	for r1 != nil && r2 != nil {
		if r1.Val < r2.Val {
			s.Next = r1
			s = r1
			r1 = r1.Next
		} else {
			s.Next = r2
			s = r2
			r2 = r2.Next
		}
	}
	if r1 != nil {
		s.Next = r1

	} else {
		s.Next = r2
	}
	L0 = L0.Next
	return L0
}
