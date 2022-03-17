package common

import (
	"fmt"
	"testing"
)

func TestArray2List(t *testing.T) {
	l := Array2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(List2Array(l))
}
