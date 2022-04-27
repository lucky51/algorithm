package datastructure

import (
	"fmt"
	"testing"
)

func TestCountRectangles(t *testing.T) {
	countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{2, 1}, {1, 4}})
}

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
