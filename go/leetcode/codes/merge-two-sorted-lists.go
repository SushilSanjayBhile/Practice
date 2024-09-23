package codes

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createTwoLists(list1 []int) *ListNode {
	start := ListNode{0, nil}
	temp := &ListNode{}

	for idx, val := range list1 {
		if idx == 0 {
			start = ListNode{val, nil}
			temp = &start
		} else {
			node := ListNode{val, nil}
			temp.Next = &node
			temp = &node
		}
	}

	return &start
}

func printList(list1 *ListNode) {
	fmt.Println()
	for {
		fmt.Print(list1.Val, " ")
		list1 = list1.Next
		if list1 == nil {
			break
		}
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return result.Next
}

func MergeTwoLists() {
	list1data := []int{1, 2, 4, 5, 6}
	list2data := []int{1, 3, 4}

	list1 := createTwoLists(list1data)
	list2 := createTwoLists(list2data)

	printList(list1)
	printList(list2)

	finalList := mergeTwoLists(list1, list2)
	printList(finalList)
}
