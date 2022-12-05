package new_user_list

import (
	"fmt"
	"goland_code/type_def"
	"goland_code/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	length := getLen(head)
	if length == 0 {
		return false
	}

	tmpHead := head
	var reversedHead *ListNode
	midIndex := length / 2

	for i := 0; i < midIndex; i++ {
		reversedHead, tmpHead = reverseIn(reversedHead, tmpHead)
	}

	if length%2 != 0 {
		tmpHead = tmpHead.Next
	}

	return isSameLink(tmpHead, reversedHead)
}

func isSameLink(link1 *ListNode, link2 *ListNode) bool {
	for link1 != nil && link2 != nil {
		if link1.Val != link2.Val {
			return false
		}
		link1 = link1.Next
		link2 = link2.Next
	}
	if link1 == nil && link2 == nil {
		return true
	}

	return false
}

func reverseIn(reHead *ListNode, head *ListNode) (*ListNode, *ListNode) {
	if reHead == nil {
		reHead = head
		head = head.Next
		reHead.Next = nil
	} else {
		tmpNode := head
		head = head.Next
		tmpNode.Next = reHead
		reHead = tmpNode
	}

	return reHead, head
}

func getLen(head *ListNode) (length int) {
	tmpHead := head
	for tmpHead != nil {
		length++
		tmpHead = tmpHead.Next
	}
	return length
}

func InitPalindromeLinkedListFunc(funcMap map[int]type_def.LProjectImpl) {
	lP := type_def.NewLProjectImpl("回文链表", 0, runIsPalindrome, 234)
	funcMap[lP.CodeNum] = lP
}

func runIsPalindrome() {

	var head ListNode
	head.Val = 9
	head.Next = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	fmt.Printf("args |\t %v \n", head)
	results := isPalindrome(&head)
	fmt.Printf("resp |\t %v \n", results)
	utils.PrintLine()
}
