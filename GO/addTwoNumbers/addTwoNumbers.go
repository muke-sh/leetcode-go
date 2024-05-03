package addtwonumbers

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) String() string {
	listVals := ""

	for list != nil {
		listVals += fmt.Sprintf(" %d ", list.Val)

		list = list.Next
	}

	return fmt.Sprintf("[ %s ]\n", listVals)
}

func (list *ListNode) PrintList() {

	listVals := ""

	for list != nil {
		listVals += fmt.Sprintf(" %d ", list.Val)

		list = list.Next
	}

	fmt.Printf("[ %s ]\n", listVals)
}

func CreateLinkedList(nums []int) *ListNode {
	var newList *ListNode = &ListNode{}

	currentNode := newList

	for i := 0; i < len(nums); i++ {

		if i != len(nums)-1 {
			newNode := &ListNode{}
			currentNode.Val = nums[i]
			currentNode.Next = newNode

			currentNode = currentNode.Next
		} else {
			currentNode.Val = nums[i]
		}
	}

	return newList
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var solutionList *ListNode = &ListNode{}

	l1CurrentNode := l1
	l2CurrentNode := l2
	solutionListCurrentNode := solutionList

	carry := 0
	for l1CurrentNode != nil || l2CurrentNode != nil {

		l1CurVal := 0
		l2CurVal := 0

		if l1CurrentNode != nil {
			l1CurVal = l1CurrentNode.Val
		}

		if l2CurrentNode != nil {
			l2CurVal = l2CurrentNode.Val
		}

		addition := l1CurVal + l2CurVal + carry
		take := addition % 10

		if addition >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		solutionListCurrentNode.Val = take

		if (l1CurrentNode != nil && l1CurrentNode.Next != nil) ||
			(l2CurrentNode != nil && l2CurrentNode.Next != nil) {

			var newNode *ListNode = &ListNode{}
			solutionListCurrentNode.Next = newNode

			solutionListCurrentNode = solutionListCurrentNode.Next
		} else if carry != 0 {
			var newNode *ListNode = &ListNode{}

			solutionListCurrentNode.Next = newNode

			solutionListCurrentNode = solutionListCurrentNode.Next
			solutionListCurrentNode.Val = carry
		}

		if l1CurrentNode != nil {
			l1CurrentNode = l1CurrentNode.Next
		}

		if l2CurrentNode != nil {
			l2CurrentNode = l2CurrentNode.Next
		}

	}

	return solutionList
}
