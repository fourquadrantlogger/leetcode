package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func Array2List(array []int) (l *ListNode) {
	if len(array) > 0 {
		l = &ListNode{
			Val: array[0],
		}
	}
	x := l
	for _, a := range array[1:] {
		x.Next = &ListNode{
			Val: a,
		}
		x = x.Next
	}
	return
}

func List2Array(l *ListNode) (r []int) {
	x := l
	for x != nil {
		r = append(r, x.Val)
		x = x.Next
	}
	return
}
