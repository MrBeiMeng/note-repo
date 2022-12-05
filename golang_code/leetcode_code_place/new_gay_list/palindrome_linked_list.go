package new_gay_list

import "goland_code/utils"

func isPalindrome(head *utils.ListNode) bool {
	length := getLen(head)
	if length == 0 {
		return false
	}

	tmpHead := head
	var reversedHead *utils.ListNode
	midIndex := length / 2

	for i := 0; i < midIndex; i++ {
		reversedHead, tmpHead = reverseIn(reversedHead, tmpHead)
	}

	if length%2 != 0 {
		tmpHead = tmpHead.Next
	}

	return isSameLink(tmpHead, reversedHead)
}

func isSameLink(link1 *utils.ListNode, link2 *utils.ListNode) bool {
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

func reverseIn(reHead *utils.ListNode, head *utils.ListNode) (*utils.ListNode, *utils.ListNode) {
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

func getLen(head *utils.ListNode) (length int) {
	tmpHead := head
	for tmpHead != nil {
		length++
		tmpHead = tmpHead.Next
	}
	return length
}

type PalindromeLinkedList struct {
}

func (p PalindromeLinkedList) Run() {
	head := utils.GetIntListNodeHandler(1)
	print("\t参数")
	head.Print()

	print("\t结果")
	println(isPalindrome(head))
}

func (p PalindromeLinkedList) RunWithoutArgs() {
	print("\t参数")
	head := utils.GetIntListNode(1, 2, 2, 1)
	head.Print()

	print("\t结果")
	println(isPalindrome(head))
}

func (p PalindromeLinkedList) GetLevel() string {
	return utils.GetEasy()
}

func (p PalindromeLinkedList) GetStar() string {
	return utils.GetStar(2)
}

func (p PalindromeLinkedList) GetStatus() string {
	return utils.GetDone()
}

func (p PalindromeLinkedList) WillShow() bool {
	return true
}

func (p PalindromeLinkedList) GetCodeNum() int {
	return 234
}

func (p PalindromeLinkedList) GetCodeTitle() string {
	return "回文链表"
}

func (p PalindromeLinkedList) GetDescription() string {
	return "给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。"
}

func (p PalindromeLinkedList) GetArgsDescription() string {
	return "\n\t链表中节点数目在范围[1, 105] 内\n\t0 <= Node.val <= 9"
}

func (p PalindromeLinkedList) GetExamples() string {
	return "\n\t示例 1：\n\t\t输入：head = [1,2,2,1]\n\t\t输出：true\n\t示例 2：\n\t\t输入：head = [1,2]\n\t\t输出：false"
}

func (p PalindromeLinkedList) GetTags() []string {
	return []string{"栈", "递归", "链表", "双指针"}
}
func (p PalindromeLinkedList) GetUrl() string {
	return "https://leetcode.cn/problems/palindrome-linked-list/"
}
