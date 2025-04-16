package linkedlist

import (
    "testing"
)

// createList 辅助函数：根据数组创建链表
func createList(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }
    
    head := &ListNode{Val: nums[0]}
    current := head
    for i := 1; i < len(nums); i++ {
        current.Next = &ListNode{Val: nums[i]}
        current = current.Next
    }
    return head
}

// compareList 辅助函数：比较两个链表是否相同
func compareList(l1, l2 *ListNode) bool {
    for l1 != nil && l2 != nil {
        if l1.Val != l2.Val {
            return false
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    return l1 == nil && l2 == nil
}

func TestMergeTwoLists(t *testing.T) {
    tests := []struct {
        name     string
        list1    []int
        list2    []int
        expected []int
    }{
        {
            name:     "两个非空链表",
            list1:    []int{1, 2, 4},
            list2:    []int{1, 3, 4},
            expected: []int{1, 1, 2, 3, 4, 4},
        },
        {
            name:     "第一个链表为空",
            list1:    []int{},
            list2:    []int{1, 2, 3},
            expected: []int{1, 2, 3},
        },
        {
            name:     "第二个链表为空",
            list1:    []int{1, 2, 3},
            list2:    []int{},
            expected: []int{1, 2, 3},
        },
        {
            name:     "两个空链表",
            list1:    []int{},
            list2:    []int{},
            expected: []int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l1 := createList(tt.list1)
            l2 := createList(tt.list2)
            expected := createList(tt.expected)
            result := MergeTwoLists(l1, l2)
            
            if !compareList(result, expected) {
                t.Errorf("测试失败：%s", tt.name)
            }
        })
    }
}