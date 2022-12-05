package utils

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	tmpNode := l
	answerNums := make([]string, 0)
	for tmpNode != nil {
		answerNums = append(answerNums, strconv.Itoa(tmpNode.Val))
		tmpNode = tmpNode.Next
	}

	pStr := strings.Join(answerNums, ",")
	fmt.Printf("[%s]\t", pStr)
}

// GetIntListNodeHandler 获取链表
//
//	手动输入
func GetIntListNodeHandler(argNum int) (head *ListNode) {
	numsStr, err := GetInput(fmt.Sprintf("\t参数序号%d\t请输入一个链表,例如:[1,2,3,4]", argNum), 1)
	if err != nil {
		println(err.Error())
		syscall.Exit(-1)
	}

	if strings.Contains(numsStr, "[") {
		numsStr = numsStr[1:]
	}
	if strings.Contains(numsStr, "]") {
		numsStr = numsStr[:len(numsStr)-1]
	}

	numStrArr := strings.Split(numsStr, ",")

	numIntArr := make([]int, 0)
	for _, s := range numStrArr {
		num, _ := strconv.Atoi(s)
		numIntArr = append(numIntArr, num)
	}

	return GetIntListNode(numIntArr...)
}

// GetIntListNode 获取链表
//
//	通过参数
func GetIntListNode(nums ...int) (head *ListNode) {
	var tmpHead *ListNode
	tmpHead = &ListNode{}
	head = tmpHead

	for i := 0; i < len(nums); i++ {
		tmpHead.Val = nums[i]
		if i+1 < len(nums) {
			nextHead := ListNode{}
			tmpHead.Next = &nextHead
			tmpHead = tmpHead.Next
		}
	}

	return head
}

// 获取数组

// 获取字符串
