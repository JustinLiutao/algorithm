package linkedlist

/**
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 *
 * 示例 1：
 * 输入：l1 = [1,2,4], l2 = [1,3,4]
 * 输出：[1,1,2,3,4,4]
 *
 * 示例 2：
 * 输入：l1 = [], l2 = []
 * 输出：[]
 *
 * 示例 3：
 * 输入：l1 = [], l2 = [0]
 * 输出：[0]
 *
 * 提示：
 * 两个链表的节点数目范围是 [0, 50]
 * -100 <= Node.val <= 100
 * l1 和 l2 均按 非递减顺序 排列
*/


// ListNode 定义链表节点结构
type ListNode struct {
    Val  int
    Next *ListNode
}

// MergeTwoLists 合并两个有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    // 创建哨兵节点
    dummy := &ListNode{Val: 0}
    current := dummy
    
    // 同时遍历两个链表
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }
    
    // 处理剩余节点
    if l1 != nil {
        current.Next = l1
    }
    if l2 != nil {
        current.Next = l2
    }
    
    return dummy.Next
}