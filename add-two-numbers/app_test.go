package add_two_numbers

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	a, b := common.Array2List([]int{1, 2, 3}), common.Array2List([]int{3, 2, 1})
	c := AddTwoNumbers(a, b)
	fmt.Println(common.List2Array(c))
}
